/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { comment, KnownMediaType } from '@azure-tools/codegen';
import { ArraySchema, ByteArraySchema, ChoiceSchema, CodeModel, ConstantSchema, DateTimeSchema, DictionarySchema, GroupProperty, ImplementationLocation, NumberSchema, ObjectSchema, Operation, Parameter, Property, Protocols, Response, Schema, SchemaResponse, SchemaType } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { aggregateParameters, getSchemaResponse, isArraySchema, isMultiRespOperation, isPageableOperation, isSchemaResponse, isTypePassedByValue, PagerInfo, isLROOperation, commentLength } from '../common/helpers';
import { OperationNaming } from '../transform/namer';
import { contentPreamble, emitPoller, formatParameterTypeName, formatStatusCodes, getClientPipeline, getFinalResponseEnvelopeName, getResponseEnvelope, getResponseEnvelopeName, getResultFieldName, getStatusCodes, hasDescription, hasResultEnvelope, hasSchemaResponse, skipURLEncoding, sortAscending, getCreateRequestParameters, getCreateRequestParametersSig, getMethodParameters, getParamName, formatParamValue, dateFormat, datetimeRFC1123Format, datetimeRFC3339Format, sortParametersByRequired } from './helpers';
import { ImportManager } from './imports';

// represents the generated content for an operation group
export class OperationGroupContent {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for all <operation>.go files
export async function generateOperations(session: Session<CodeModel>): Promise<OperationGroupContent[]> {
  const isARM = session.model.language.go!.openApiType === 'arm';
  const forceExports = <boolean>session.model.language.go!.exportClients;
  // generate protocol operations
  const operations = new Array<OperationGroupContent>();
  for (const group of values(session.model.operationGroups)) {
    // the list of packages to import
    const imports = new ImportManager();
    // add standard imports
    imports.add('net/http');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/policy');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    if (<boolean>session.model.language.go!.azureARM) {
      imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm');
    }

    let opText = '';
    group.operations.sort((a: Operation, b: Operation) => { return sortAscending(a.language.go!.name, b.language.go!.name) });
    for (const op of values(group.operations)) {
      // protocol creation can add imports to the list so
      // it must be done before the imports are written out
      if (isLROOperation(op)) {
        // generate Begin method
        opText += generateLROBeginMethod(op, imports, isARM);
      }
      opText += generateOperation(op, imports);
      opText += createProtocolRequest(session.model, op, imports);
      if (!isLROOperation(op) || isPageableOperation(op)) {
        // LRO responses are handled elsewhere, with the exception of pageable LROs
        opText += createProtocolResponse(op, imports);
      }
      if (!op.language.go!.headAsBoolean) {
        // HEAD requests never return a response body so there's no error to unmarshal
        opText += createProtocolErrHandler(op, imports);
      }
    }
    // stitch it all together
    let text = await contentPreamble(session);
    let connection = 'Connection';
    let clientName = group.language.go!.clientName;
    if (<boolean>session.model.language.go!.azureARM) {
      connection = 'arm.Connection';
    } else if (!forceExports) {
      connection = connection.uncapitalize();
    }
    const clientCtor = group.language.go!.clientCtorName;
    text += imports.text();
    // generate the operation client
    if (isARM || forceExports) {
      text += `// ${clientName} contains the methods for the ${group.language.go!.name} group.\n`;
      text += `// Don't use this type directly, use ${clientCtor}() instead.\n`;
    }
    text += `type ${clientName} struct {\n`;
    if (<boolean>session.model.language.go!.azureARM) {
      text += '\tep string\n';
      text += '\tpl runtime.Pipeline\n';
    } else {
      text += `\tcon *${connection}\n`;
    }
    if (group.language.go!.clientParams) {
      const clientParams = <Array<Parameter>>group.language.go!.clientParams;
      for (const clientParam of values(clientParams)) {
        text += `\t${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}\n`;
      }
    }
    text += '}\n\n';
    if (isARM || forceExports) {
      const connectionLiterals = new Array<string>();
      const methodParams = new Array<string>();
      if (<boolean>session.model.language.go!.azureARM) {
        // real ARM doesn't need to deal with parameterized host (yay)
        connectionLiterals.push('ep: con.Endpoint()');
        connectionLiterals.push('pl: con.NewPipeline(module, version)');
        methodParams.push('con *arm.Connection');
      } else {
        // operation client constructor
        connectionLiterals.push('con: con');
        methodParams.push(`con *${connection}`);
      }
      // add client params to the operation client constructor
      if (group.language.go!.clientParams) {
        const clientParams = <Array<Parameter>>group.language.go!.clientParams;
        clientParams.sort(sortParametersByRequired);
        for (const clientParam of values(clientParams)) {
          connectionLiterals.push(`${clientParam.language.go!.name}: ${clientParam.language.go!.name}`);
          methodParams.push(`${clientParam.language.go!.name} ${formatParameterTypeName(clientParam)}`);
        }
      }
      text += `// ${clientCtor} creates a new instance of ${clientName} with the specified values.\n`;
      text += `func ${clientCtor}(${methodParams.join(', ')}) *${clientName} {\n`;
      text += `\treturn &${clientName}{${connectionLiterals.join(', ')}}\n`;
      text += '}\n\n';
    }
    // add operations content last
    text += opText;
    operations.push(new OperationGroupContent(group.language.go!.name, text));
  }
  return operations;
}

// use this to generate the code that will help process values returned in response headers
function formatHeaderResponseValue(propName: string, header: string, schema: Schema, imports: ImportManager, respObj: string, zeroResp: string): string {
  // dictionaries are handled slightly different so we do that first
  if (schema.type === SchemaType.Dictionary) {
    imports.add('strings');
    const headerPrefix = schema.language.go!.headerCollectionPrefix;
    let text = '\tfor hh := range resp.Header {\n';
    text += `\t\tif len(hh) > len("${headerPrefix}") && strings.EqualFold(hh[:len("${headerPrefix}")], "${headerPrefix}") {\n`;
    text += `\t\t\tif ${respObj}.Metadata == nil {\n`;
    text += `\t\t\t\t${respObj}.Metadata = map[string]string{}\n`;
    text += '\t\t\t}\n';
    text += `\t\t\t${respObj}.Metadata[hh[len("${headerPrefix}"):]] = resp.Header.Get(hh)\n`;
    text += '\t\t}\n';
    text += '\t}\n';
    return text;
  }
  let text = `\tif val := resp.Header.Get("${header}"); val != "" {\n`;
  const name = propName.uncapitalize();
  let byRef = '&';
  switch (schema.type) {
    case SchemaType.Boolean:
      imports.add('strconv');
      text += `\t\t${name}, err := strconv.ParseBool(val)\n`;
      break;
    case SchemaType.ByteArray:
      // ByteArray is a base-64 encoded value in string format
      imports.add('encoding/base64');
      let byteFormat = 'Std';
      if ((<ByteArraySchema>schema).format === 'base64url') {
        byteFormat = 'RawURL';
      }
      text += `\t\t${name}, err := base64.${byteFormat}Encoding.DecodeString(val)\n`;
      byRef = '';
      break;
    case SchemaType.Choice:
    case SchemaType.SealedChoice:
      text += `\t\t${respObj}.${propName} = (*${schema.language.go!.name})(&val)\n`;
      text += '\t}\n';
      return text;
    case SchemaType.Constant:
    case SchemaType.Duration:
    case SchemaType.String:
      text += `\t\t${respObj}.${propName} = &val\n`;
      text += '\t}\n';
      return text;
    case SchemaType.Date:
      imports.add('time');
      text += `\t\t${name}, err := time.Parse("${dateFormat}", val)\n`;
      break;
    case SchemaType.DateTime:
      imports.add('time');
      let format = datetimeRFC3339Format;
      const dateTime = <DateTimeSchema>schema;
      if (dateTime.format === 'date-time-rfc1123') {
        format = datetimeRFC1123Format;
      }
      text += `\t\t${name}, err := time.Parse(${format}, val)\n`;
      break;
    case SchemaType.Integer:
      imports.add('strconv');
      const intNum = <NumberSchema>schema;
      if (intNum.precision === 32) {
        text += `\t\t${name}32, err := strconv.ParseInt(val, 10, 32)\n`;
        text += `\t\t${name} := int32(${name}32)\n`;
      } else {
        text += `\t\t${name}, err := strconv.ParseInt(val, 10, 64)\n`;
      }
      break;
    case SchemaType.Number:
      imports.add('strconv');
      const floatNum = <NumberSchema>schema;
      if (floatNum.precision === 32) {
        text += `\t\t${name}32, err := strconv.ParseFloat(val, 32)\n`;
        text += `\t\t${name} := float32(${name}32)\n`;
      } else {
        text += `\t\t${name}, err := strconv.ParseFloat(val, 64)\n`;
      }
      break;
    default:
      throw new Error(`unsupported header type ${schema.type}`);
  }
  text += `\t\tif err != nil {\n`;
  text += `\t\t\treturn ${zeroResp}, err\n`;
  text += `\t\t}\n`;
  text += `\t\t${respObj}.${propName} = ${byRef}${name}\n`;
  text += '\t}\n';
  return text;
}

function getZeroReturnValue(op: Operation, apiType: 'api' | 'op' | 'handler'): string {
  let returnType = `${getResponseEnvelopeName(op)}{}`;
  if (isLROOperation(op)) {
    if (apiType === 'op') {
      // the operation returns an *http.Response
      returnType = 'nil';
    } else if (apiType === 'handler' && isPageableOperation(op)) {
      returnType = `${getFinalResponseEnvelopeName(op)}{}`;
    }
  }
  return returnType
}

// returns true if the response contains any headers
function responseHasHeaders(op: Operation): boolean {
  const resultEnv = hasResultEnvelope(op);
  if (!resultEnv) {
    return false;
  }
  for (const prop of values((<ObjectSchema>resultEnv.schema).properties)) {
    if (prop.language.go!.fromHeader) {
      return true;
    }
  }
  return false;
}

function generateOperation(op: Operation, imports: ImportManager): string {
  if (op.language.go!.paging && op.language.go!.paging.isNextOp) {
    // don't generate a public API for the methods used to advance pages
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, 'op');
  const clientName = op.language.go!.clientName;
  let text = '';
  if (hasDescription(op.language.go!)) {
    text += `${comment(`${op.language.go!.name} - ${op.language.go!.description}`, "//", undefined, commentLength)}\n`;
  }
  let opName = op.language.go!.name;
  if (isLROOperation(op)) {
    opName = info.protocolNaming.internalMethod;
  }
  text += `func (client *${clientName}) ${opName}(${params}) (${returns.join(', ')}) {\n`;
  const reqParams = getCreateRequestParameters(op);
  const statusCodes = getStatusCodes(op);
  if (isPageableOperation(op) && !isLROOperation(op)) {
    imports.add('context');
    text += `\treturn &${(<PagerInfo>op.language.go!.pageableType).name}{\n`;
    text += `\t\tclient: client,\n`;
    text += `\t\trequester: func(ctx context.Context) (*policy.Request, error) {\n`;
    text += `\t\t\treturn client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
    text += '\t\t},\n';
    text += `\t\tadvancer: func(ctx context.Context, resp ${getResponseEnvelopeName(op)}) (*policy.Request, error) {\n`;
    const nextLink = op.language.go!.paging.nextLinkName;
    const response = getResultFieldName(op);
    if (op.language.go!.paging.member) {
      const nextOpParams = getCreateRequestParametersSig(op.language.go!.paging.nextLinkOperation).split(',');
      // keep the parameter names from the name/type tuples and find nextLink param
      for (let i = 0; i < nextOpParams.length; ++i) {
        const paramName = nextOpParams[i].trim().split(' ')[0];
        const paramType = nextOpParams[i].trim().split(' ')[1];
        if (paramName.startsWith('next') && paramType === 'string') {
          nextOpParams[i] = `*resp.${response}.${nextLink}`;
        } else {
          nextOpParams[i] = paramName;
        }
      }
      text += `\t\t\treturn client.${op.language.go!.paging.member}CreateRequest(${nextOpParams.join(', ')})\n`;
    } else {
      text += `\t\t\treturn runtime.NewRequest(ctx, http.MethodGet, *resp.${response}.${nextLink})\n`;
    }
    text += '\t\t},\n';
    text += `\t}\n`;
    text += '}\n\n';
    return text;
  }
  const zeroResp = getZeroReturnValue(op, 'op');
  text += `\treq, err := client.${info.protocolNaming.requestMethod}(${reqParams})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn ${zeroResp}, err\n`;
  text += `\t}\n`;
  text += `\tresp, err := ${getClientPipeline(op)}.Do(req)\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn ${zeroResp}, err\n`;
  text += `\t}\n`;
  // HAB with headers response is handled in protocol responder
  if (op.language.go!.headAsBoolean && !responseHasHeaders(op)) {
    text += `\tresult := ${getResponseEnvelopeName(op)}{RawResponse: resp}\n`;
    text += '\tif resp.StatusCode >= 200 && resp.StatusCode < 300 {\n';
    text += '\t\tresult.Success = true\n';
    text += '\t}\n';
    text += '\treturn result, nil\n';
  } else {
    // for complex HAB the status code check isn't applicable
    if (!op.language.go!.headAsBoolean) {
      text += `\tif !runtime.HasStatusCode(resp, ${formatStatusCodes(statusCodes)}) {\n`;
      text += `\t\treturn ${zeroResp}, client.${info.protocolNaming.errorMethod}(resp)\n`;
      text += '\t}\n';
    }
    if (isLROOperation(op)) {
      text += '\t return resp, nil\n';
    } else if (needsResponseHandler(op)) {
      // also cheating here as at present the only param to the responder is an http.Response
      text += `\treturn client.${info.protocolNaming.responseMethod}(resp)\n`;
    } else {
      text += `\treturn ${getResponseEnvelopeName(op)}{RawResponse: resp}, nil\n`;
    }
  }
  text += '}\n\n';
  return text;
}

function createProtocolRequest(codeModel: CodeModel, op: Operation, imports: ImportManager): string {
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.requestMethod;
  for (const param of values(aggregateParameters(op))) {
    if (param.implementation !== ImplementationLocation.Method || param.required !== true) {
      continue;
    }
    imports.addImportForSchemaType(param.schema);
  }
  const returns = ['*policy.Request', 'error'];
  let text = `${comment(name, '// ')} creates the ${info.name} request.\n`;
  text += `func (client *${op.language.go!.clientName}) ${name}(${getCreateRequestParametersSig(op)}) (${returns.join(', ')}) {\n`;
  // default to host on the connection
  let hostParam = 'client.con.Endpoint()';
  if (<boolean>op.language.go!.azureARM) {
    hostParam = '\tclient.ep';
  }
  if (codeModel.language.go!.complexHostParams) {
    imports.add('strings');
    // we have a complex parameterized host
    text += `\thost := "${op.requests![0].protocol.http!.uri}"\n`;
    // get all the host params on the connection
    const hostParams = <Array<Parameter>>codeModel.language.go!.hostParams;
    for (const hostParam of values(hostParams)) {
      text += `\thost = strings.ReplaceAll(host, "{${hostParam.language.go!.serializedName}}", client.con.${(<string>hostParam.language.go!.name).capitalize()}())\n`;
    }
    // check for any method local host params
    for (const param of values(op.parameters)) {
      if (param.implementation === ImplementationLocation.Method && param.protocol.http!.in === 'uri') {
        text += `\thost = strings.ReplaceAll(host, "{${param.language.go!.serializedName}}", ${param.language.go!.name})\n`;
      }
    }
    hostParam = 'host';
  }
  const hasPathParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; }).any();
  // storage needs the client.u to be the source-of-truth for the full path.
  // however, swagger requires that all operations specify a path, which is at odds with storage.
  // to work around this, storage specifies x-ms-path paths with path params but doesn't
  // actually reference the path params (i.e. no params with which to replace the tokens).
  // so, if a path contains tokens but there are no path params, skip emitting the path.
  const pathStr = <string>op.requests![0].protocol.http!.path;
  const pathContainsParms = pathStr.includes('{');
  if (hasPathParams || (!pathContainsParms && pathStr.length > 1)) {
    // there are path params, or the path doesn't contain tokens and is not "/" so emit it
    text += `\turlPath := "${op.requests![0].protocol.http!.path}"\n`;
    hostParam = `runtime.JoinPaths(${hostParam}, urlPath)`;
  }
  if (hasPathParams) {
    // swagger defines path params, emit path and replace tokens
    imports.add('strings');
    // replace path parameters
    for (const pp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'path'; })) {
      // emit check to ensure path param isn't an empty string.  we only need
      // to do this for params that have an underlying type of string.
      const choiceIsString = function (schema: Schema): boolean {
        if (schema.type === SchemaType.Choice) {
          return (<ChoiceSchema>schema).choiceType.type === SchemaType.String;
        }
        if (schema.type === SchemaType.SealedChoice) {
          return (<ChoiceSchema>schema).choiceType.type === SchemaType.String;
        }
        return false;
      }
      if (pp.schema.type === SchemaType.String || choiceIsString(pp.schema)) {
        const paramName = getParamName(pp);
        imports.add('errors');
        text += `\tif ${paramName} == "" {\n`;
        text += `\t\treturn nil, errors.New("parameter ${paramName} cannot be empty")\n`;
        text += '\t}\n';
      }
      let paramValue = formatParamValue(pp, imports);
      if (!skipURLEncoding(pp)) {
        imports.add('net/url');
        paramValue = `url.PathEscape(${formatParamValue(pp, imports)})`;
      }
      text += `\turlPath = strings.ReplaceAll(urlPath, "{${pp.language.go!.serializedName}}", ${paramValue})\n`;
    }
  }
  text += `\treq, err := runtime.NewRequest(ctx, http.Method${(<string>op.requests![0].protocol.http!.method).capitalize()}, ${hostParam})\n`;
  text += '\tif err != nil {\n';
  text += '\t\treturn nil, err\n';
  text += '\t}\n';
  const hasQueryParams = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; }).any();
  // helper to build nil checks for param groups
  const emitParamGroupCheck = function (gp: GroupProperty, param: Parameter): string {
    if (param.implementation === ImplementationLocation.Client) {
      return `\tif client.${param.language.go!.name} != nil {\n`;
    }
    const paramGroupName = (<string>gp.language.go!.name).uncapitalize();
    let optionalParamGroupCheck = `${paramGroupName} != nil && `;
    if (gp.required) {
      optionalParamGroupCheck = '';
    }
    return `\tif ${optionalParamGroupCheck}${paramGroupName}.${(<string>param.language.go!.name).capitalize()} != nil {\n`;
  }
  if (hasQueryParams) {
    // add query parameters
    const encodedParams = new Array<Parameter>();
    const unencodedParams = new Array<Parameter>();
    for (const qp of values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined && each.protocol.http!.in === 'query'; })) {
      if (skipURLEncoding(qp)) {
        unencodedParams.push(qp);
      } else {
        encodedParams.push(qp);
      }
    }
    const emitQueryParam = function (qp: Parameter, setter: string): string {
      let qpText = '';
      if (qp.required === true) {
        qpText = `\t${setter}\n`;
      } else if (qp.implementation === ImplementationLocation.Client) {
        // global optional param
        qpText = `\tif client.${qp.language.go!.name} != nil {\n`;
        qpText += `\t\t${setter}\n`;
        qpText += `\t}\n`;
      } else {
        qpText = emitParamGroupCheck(<GroupProperty>qp.language.go!.paramGroup, qp);
        qpText += `\t\t${setter}\n`;
        qpText += `\t}\n`;
      }
      return qpText;
    }
    // emit encoded params first
    if (encodedParams.length > 0) {
      text += '\treqQP := req.Raw().URL.Query()\n';
      for (const qp of values(encodedParams)) {
        let setter: string;
        if (qp.protocol.http?.explode === true) {
          setter = `\tfor _, qv := range ${getParamName(qp)} {\n`;
          if (qp.schema.type !== SchemaType.Array) {
            throw new Error(`expected SchemaType.Array for query param ${qp.language.go!.name}`);
          }
          // emit a type conversion for the qv based on the array's element type
          let queryVal: string;
          const arrayQP = <ArraySchema>qp.schema;
          switch (arrayQP.elementType.type) {
            case SchemaType.Choice:
            case SchemaType.SealedChoice:
              const ch = <ChoiceSchema>arrayQP.elementType;
              // only string and number types are supported for enums
              if (ch.choiceType.type === SchemaType.String) {
                queryVal = 'string(qv)';
              } else {
                imports.add('fmt');
                queryVal = 'fmt.Sprintf("%d", qv)';
              }
              break;
            case SchemaType.String:
              queryVal = 'qv';
              break;
            default:
              imports.add('fmt');
              queryVal = 'fmt.Sprintf("%v", qv)';
          }
          setter += `\t\treqQP.Add("${qp.language.go!.serializedName}", ${queryVal})\n`;
          setter += '\t}';
        } else {
          // cannot initialize setter to this value as formatParamValue() can change imports
          setter = `reqQP.Set("${qp.language.go!.serializedName}", ${formatParamValue(qp, imports)})`;
        }
        text += emitQueryParam(qp, setter);
      }
      text += '\treq.Raw().URL.RawQuery = reqQP.Encode()\n';
    }
    // tack on any unencoded params to the end
    if (unencodedParams.length > 0) {
      if (encodedParams.length > 0) {
        text += '\tunencodedParams := []string{req.Raw().URL.RawQuery}\n';
      } else {
        text += '\tunencodedParams := []string{}\n';
      }
      for (const qp of values(unencodedParams)) {
        let setter: string;
        if (qp.protocol.http?.explode === true) {
          setter = `\tfor _, qv := range ${getParamName(qp)} {\n`;
          setter += `\t\tunencodedParams = append(unencodedParams, "${qp.language.go!.serializedName}="+qv)\n`;
          setter += '\t}';
        } else {
          setter = `unencodedParams = append(unencodedParams, "${qp.language.go!.serializedName}="+${formatParamValue(qp, imports)})`;
        }
        text += emitQueryParam(qp, setter);
      }
      text += '\treq.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")\n';
    }
  }
  if (hasBinaryResponse(op.responses!)) {
    // skip auto-body downloading for binary stream responses
    text += '\treq.SkipBodyDownload()\n';
  }
  // add specific request headers
  const headerParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http !== undefined; }).where((each: Parameter) => { return each.protocol.http!.in === 'header'; });
  headerParam.forEach(header => {
    const emitHeaderSet = function (headerParam: Parameter, prefix: string): string {
      if (header.schema.language.go!.headerCollectionPrefix) {
        let headerText = `${prefix}for k, v := range ${getParamName(headerParam)} {\n`;
        headerText += `${prefix}\treq.Raw().Header.Set("${header.schema.language.go!.headerCollectionPrefix}"+k, v)\n`;
        headerText += `${prefix}}\n`;
        return headerText;
      } else {
        return `${prefix}req.Raw().Header.Set("${headerParam.language.go!.serializedName}", ${formatParamValue(headerParam, imports)})\n`;
      }
    }
    if (header.required) {
      text += emitHeaderSet(header, '\t');
    } else {
      text += emitParamGroupCheck(<GroupProperty>header.language.go!.paramGroup, header);
      text += emitHeaderSet(header, '\t\t');
      text += `\t}\n`;
    }
  });
  const mediaType = getMediaType(op.requests![0].protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    // default to the body param name
    let body = getParamName(bodyParam!);
    if (bodyParam!.schema.type === SchemaType.Constant) {
      // if the value is constant, embed it directly
      body = formatConstantValue(<ConstantSchema>bodyParam!.schema);
    } else if (mediaType === 'XML' && bodyParam!.schema.type === SchemaType.Array) {
      // for XML payloads, create a wrapper type if the payload is an array
      imports.add('encoding/xml');
      text += '\ttype wrapper struct {\n';
      let tagName = bodyParam!.schema.language.go!.name;
      if (bodyParam!.schema.serialization?.xml?.name) {
        tagName = bodyParam!.schema.serialization.xml.name;
      }
      text += `\t\tXMLName xml.Name \`xml:"${tagName}"\`\n`;
      let fieldName = bodyParam!.schema.language.go!.name;
      if (isArraySchema(bodyParam!.schema)) {
        fieldName = (<string>bodyParam!.language.go!.name).capitalize();
        let tag = bodyParam!.schema.elementType.language.go!.name;
        if (bodyParam!.schema.elementType.serialization?.xml?.name) {
          tag = bodyParam!.schema.elementType.serialization.xml.name;
        }
        text += `\t\t${fieldName} *${bodyParam!.schema.language.go!.name} \`xml:"${tag}"\`\n`;
      }
      text += '\t}\n';
      let addr = '&';
      if (bodyParam && (!bodyParam.required && !isTypePassedByValue(bodyParam.schema))) {
        addr = '';
      }
      body = `wrapper{${fieldName}: ${addr}${body}}`;
    } else if (bodyParam!.schema.type === SchemaType.Date) {
      // wrap the body in the internal dateType
      body = `dateType(${body})`;
    } else if ((bodyParam!.schema.type === SchemaType.DateTime && (<DateTimeSchema>bodyParam!.schema).format === 'date-time-rfc1123') || bodyParam!.schema.type === SchemaType.UnixTime) {
      // wrap the body in the custom RFC1123 type
      text += `\taux := ${bodyParam!.schema.language.go!.internalTimeType}(${body})\n`;
      body = 'aux';
    } else if (isArrayOfTimesForMarshalling(bodyParam!.schema) || isArrayOfDatesForMarshalling(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := make([]*${timeType}, len(${body}), len(${body}))\n`;
      text += `\tfor i := 0; i < len(${body}); i++ {\n`;
      text += `\t\taux[i] = (*${timeType})(${body}[i])\n`;
      text += '\t}\n';
      body = 'aux';
    } else if (isMapOfDateTime(bodyParam!.schema) || isMapOfDate(bodyParam!.schema)) {
      const timeType = (<ArraySchema>bodyParam!.schema).elementType.language.go!.internalTimeType;
      text += `\taux := map[string]*${timeType}{}\n`;
      text += `\tfor k, v := range ${body} {\n`;
      text += `\t\taux[k] = (*${timeType})(v)\n`;
      text += '\t}\n';
      body = 'aux';
    }
    if (bodyParam!.required || bodyParam!.schema.type === SchemaType.Constant) {
      text += `\treturn req, runtime.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, `req, ${body}`)}\n`;
    } else {
      text += emitParamGroupCheck(<GroupProperty>bodyParam!.language.go!.paramGroup, bodyParam!);
      text += `\t\treturn req, runtime.MarshalAs${getMediaFormat(bodyParam!.schema, mediaType, `req, ${body}`)}\n`;
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'binary') {
    let contentType = `"${op.requests![0].protocol.http!.mediaTypes[0]}"`;
    if (op.requests![0].protocol.http!.mediaTypes.length > 1) {
      for (const param of values(op.requests![0].parameters)) {
        // If a request defined more than one possible media type, then the param is expected to be synthesized from modelerfour
        // and should be a SealedChoice schema type that account for the acceptable media types defined in the swagger. 
        if (param.origin === 'modelerfour:synthesized/content-type' && param.schema.type === SchemaType.SealedChoice) {
          contentType = `string(${param.language.go!.name})`;
        }
      }
    }
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    if (bodyParam!.required) {
      text += `\treturn req, req.SetBody(${bodyParam?.language.go!.name}, ${contentType})\n`;
    } else {
      text += emitParamGroupCheck(<GroupProperty>bodyParam!.language.go!.paramGroup, bodyParam!);
      text += `\treturn req, req.SetBody(${getParamName(bodyParam!)}, ${contentType})\n`;
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'text') {
    imports.add('strings');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming');
    const bodyParam = values(aggregateParameters(op)).where((each: Parameter) => { return each.protocol.http!.in === 'body'; }).first();
    if (bodyParam!.required) {
      text += `\tbody := streaming.NopCloser(strings.NewReader(${bodyParam!.language.go!.name}))\n`;
      text += `\treturn req, req.SetBody(body, "text/plain; encoding=UTF-8")\n`;
    } else {
      text += emitParamGroupCheck(<GroupProperty>bodyParam!.language.go!.paramGroup, bodyParam!);
      text += `\tbody := streaming.NopCloser(strings.NewReader(${getParamName(bodyParam!)}))\n`;
      text += `\treturn req, req.SetBody(body, "text/plain; encoding=UTF-8")\n`;
      text += '\t}\n';
      text += '\treturn req, nil\n';
    }
  } else if (mediaType === 'multipart') {
    text += '\tif err := runtime.SetMultipartFormData(req, map[string]interface{}{\n';
    for (const param of values(aggregateParameters(op))) {
      if (param.isPartialBody) {
        text += `\t\t\t"${param.language.go!.name}": ${param.language.go!.name},\n`;
      }
    }
    text += '}); err != nil {'
    text += '\t\treturn nil, err\n';
    text += '\t}\n';
    text += '\treturn req, nil\n';
  } else {
    text += `\treturn req, nil\n`;
  }
  text += '}\n\n';
  return text;
}

function getMediaFormat(schema: Schema, mediaType: 'JSON' | 'XML', param: string): string {
  let marshaller: 'JSON' | 'XML' | 'ByteArray' = mediaType;
  let format = '';
  if (schema.type === SchemaType.ByteArray) {
    marshaller = 'ByteArray';
    format = ', runtime.Base64StdFormat';
    if ((<ByteArraySchema>schema).format === 'base64url') {
      format = ', runtime.Base64URLFormat';
    }
  }
  return `${marshaller}(${param}${format})`;
}

function isArrayOfTimesForMarshalling(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  if (arrayElem.type === SchemaType.UnixTime) {
    return true;
  }
  if (arrayElem.type !== SchemaType.DateTime) {
    return false;
  }
  return (<DateTimeSchema>arrayElem).format === 'date-time-rfc1123';
}

function isArrayOfDatesForMarshalling(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.Date;
}

function needsResponseHandler(op: Operation): boolean {
  return hasSchemaResponse(op) || responseHasHeaders(op) || (isLROOperation(op) && hasResultEnvelope(op) !== undefined) || isPageableOperation(op);
}

function generateResponseUnmarshaller(op: Operation, response: SchemaResponse, unmarshalTarget: string): string {
  let unmarshallerText = '';
  const zeroValue = getZeroReturnValue(op, 'handler');
  if (response.schema.type === SchemaType.DateTime || response.schema.type === SchemaType.UnixTime || response.schema.type === SchemaType.Date) {
    // use the designated time type for unmarshalling
    unmarshallerText += `\tvar aux *${response.schema.language.go!.internalTimeType}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaType(response.protocol)}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${getResultFieldName(op)} = (*time.Time)(aux)\n`;
    return unmarshallerText;
  } else if (isArrayOfDateTime(response.schema) || isArrayOfDate(response.schema)) {
    // unmarshalling arrays of date/time is a little more involved
    unmarshallerText += `\tvar aux []*${(<ArraySchema>response.schema).elementType.language.go!.internalTimeType}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaType(response.protocol)}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += '\tcp := make([]*time.Time, len(aux), len(aux))\n';
    unmarshallerText += '\tfor i := 0; i < len(aux); i++ {\n';
    unmarshallerText += '\t\tcp[i] = (*time.Time)(aux[i])\n';
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tresult.${getResultFieldName(op)} = cp\n`;
    return unmarshallerText;
  } else if (isMapOfDateTime(response.schema) || isMapOfDate(response.schema)) {
    unmarshallerText += `\taux := map[string]*${(<DictionarySchema>response.schema).elementType.language.go!.internalTimeType}{}\n`;
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaType(response.protocol)}(resp, &aux); err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
    unmarshallerText += `\tcp := map[string]*time.Time{}\n`;
    unmarshallerText += `\tfor k, v := range aux {\n`;
    unmarshallerText += `\t\tcp[k] = (*time.Time)(v)\n`;
    unmarshallerText += `\t}\n`;
    unmarshallerText += `\tresult.${getResultFieldName(op)} = cp\n`;
    return unmarshallerText;
  }
  const mediaType = getMediaType(response.protocol);
  if (mediaType === 'JSON' || mediaType === 'XML') {
    unmarshallerText += `\tif err := runtime.UnmarshalAs${getMediaFormat(response.schema, mediaType, `resp, &${unmarshalTarget}`)}; err != nil {\n`;
    unmarshallerText += `\t\treturn ${zeroValue}, err\n`;
    unmarshallerText += '\t}\n';
  }
  return unmarshallerText;
}

function createProtocolResponse(op: Operation, imports: ImportManager): string {
  if (!needsResponseHandler(op)) {
    return '';
  }
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.responseMethod;
  const clientName = op.language.go!.clientName;
  let text = `${comment(name, '// ')} handles the ${info.name} response.\n`;
  text += `func (client *${clientName}) ${name}(resp *http.Response) (${generateReturnsInfo(op, 'handler').join(', ')}) {\n`;
  const addHeaders = function (props?: Property[]) {
    const headerVals = new Array<Property>();
    for (const prop of values(props)) {
      if (prop.language.go!.fromHeader) {
        headerVals.push(prop);
      }
    }
    for (const headerVal of values(headerVals)) {
      text += formatHeaderResponseValue(headerVal.language.go!.name, headerVal.language.go!.fromHeader, headerVal.schema, imports, 'result', `${getResponseEnvelopeName(op)}{}`);
    }
  }
  if (!isMultiRespOperation(op)) {
    let respEnv = getResponseEnvelopeName(op);
    if (isLROOperation(op)) {
      respEnv = getFinalResponseEnvelopeName(op);
    }
    text += `\tresult := ${respEnv}{RawResponse: resp}\n`;
    // we know there's a result envelope at this point
    const resultEnv = hasResultEnvelope(op);
    addHeaders((<ObjectSchema>resultEnv!.schema).properties);
    const schemaResponse = getSchemaResponse(op);
    if (op.language.go!.headAsBoolean === true) {
      text += '\tif resp.StatusCode >= 200 && resp.StatusCode < 300 {\n';
      text += '\t\tresult.Success = true\n';
      text += '\t}\n';
    } else if (schemaResponse) {
      // when unmarshalling a wrapped XML array or discriminated type, unmarshal into the response envelope
      let target = `result.${getResultFieldName(op)}`
      if ((getMediaType(schemaResponse.protocol) === 'XML' && schemaResponse.schema.type === SchemaType.Array) || schemaResponse.schema.language.go!.discriminatorInterface) {
        target = 'result';
      }
      text += generateResponseUnmarshaller(op, schemaResponse, target);
    }
    text += '\treturn result, nil\n';
  } else {
    imports.add('fmt');
    text += `\tresult := ${getResponseEnvelopeName(op)}{RawResponse: resp}\n`;
    // unmarshal any header values
    const respEnv = getResponseEnvelope(op);
    addHeaders(respEnv.properties);
    text += '\tswitch resp.StatusCode {\n';
    for (const response of values(op.responses)) {
      text += `\tcase ${formatStatusCodes(response.protocol.http!.statusCodes)}:\n`
      if (!isSchemaResponse(response)) {
        // the operation contains a mix of schemas and non-schema responses
        continue;
      }
      text += `\tvar val ${response.schema.language.go!.name}\n`;
      text += generateResponseUnmarshaller(op, response, 'val');
      text += `\tresult.Value = val\n`;
    }
    text += '\tdefault:\n';
    text += `\t\treturn ${getZeroReturnValue(op, 'handler')}, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)\n`;
    text += '\t}\n';
    text += '\treturn result, nil\n';
  }
  text += '}\n\n';
  return text;
}

function createProtocolErrHandler(op: Operation, imports: ImportManager): string {
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
  const info = <OperationNaming>op.language.go!;
  const name = info.protocolNaming.errorMethod;
  let text = `${comment(name, '// ')} handles the ${info.name} error response.\n`;
  text += `func (client *${op.language.go!.clientName}) ${name}(resp *http.Response) error {\n`;
  text += '\tbody, err := runtime.Payload(resp)\n';
  text += '\tif err != nil {\n';
  text += '\t\treturn runtime.NewResponseError(err, resp)\n';
  text += '\t}\n';
  // define a generic error for when there are no exceptions or no error schema
  const generateGenericError = function () {
    imports.add('errors');
    return `\tif len(body) == 0 {
      return runtime.NewResponseError(errors.New(resp.Status), resp)
    }
    return runtime.NewResponseError(errors.New(string(body)), resp)
    `;
  }

  // if the response doesn't define any error types return a generic error
  if (!op.exceptions) {
    text += generateGenericError();
    text += '}\n\n';
    return text;
  }

  const generateUnmarshaller = function (schemaError: Schema, prefix: string) {
    let unmarshaller = '';
    if (schemaError.language.default.name === 'generic') {
      unmarshaller += `${prefix}${generateGenericError()}`;
      return unmarshaller;
    }
    const errFormat = <string>schemaError.language.go!.marshallingFormat;
    let typeName = schemaError.language.go!.name;
    if (schemaError.language.go!.internalErrorType) {
      typeName = schemaError.language.go!.internalErrorType;
    }
    imports.add('fmt');
    // for wrapped errors, raw is initialized in the unmarshaller.
    // error types other than object obviously don't have a raw field.
    if (!schemaError.language.go!.internalErrorType && schemaError.type === SchemaType.Object) {
      unmarshaller += `\t${prefix}errType := ${typeName}{raw: string(body)}\n`;
    } else {
      unmarshaller += `\tvar errType ${typeName}\n`;
    }
    const innerErr = schemaError.language.go!.flattenedErr ? `.${schemaError.language.go!.flattenedErr}` : '';
    unmarshaller += `${prefix}if err := runtime.UnmarshalAs${errFormat.toUpperCase()}(resp, &errType${innerErr}); err != nil {\n`;
    unmarshaller += `${prefix}\treturn runtime.NewResponseError(fmt.Errorf("%s\\n%s", string(body), err), resp)\n`;
    unmarshaller += `${prefix}}\n`;
    if (schemaError.language.go!.internalErrorType) {
      // err.wrapped is for discriminated error types, it will already be pointer-to-type
      unmarshaller += `${prefix}return runtime.NewResponseError(errType.wrapped, resp)\n`;
    } else if (schemaError.type === SchemaType.Object) {
      // for consistency with success responses, return pointer-to-error type
      unmarshaller += `${prefix}return runtime.NewResponseError(&errType, resp)\n`;
    } else {
      unmarshaller += `${prefix}return runtime.NewResponseError(fmt.Errorf("%v", errType), resp)\n`;
    }
    return unmarshaller;
  };
  // fold multiple error responses with the same schema into a single unmarshaller.
  const foldedMap = new Map<Schema, Array<string>>();
  // create a dummy schema for schemaless errors
  const genericErr = new Schema('generic', 'generic', SchemaType.Object);
  for (const exception of values(op.exceptions)) {
    let errSchema = genericErr;
    if (!exception.language.go!.genericError) {
      errSchema = (<SchemaResponse>exception).schema;
    }
    if (!foldedMap.has(errSchema)) {
      foldedMap.set(errSchema, new Array<string>());
    }
    for (const statusCode of values(<Array<string>>exception.protocol.http!.statusCodes)) {
      foldedMap.get(errSchema)!.push(statusCode);
    }
  }
  // only one entry in the map means all status codes return the same error schema
  if (foldedMap.size === 1) {
    text += generateUnmarshaller(values(foldedMap.keys()).first()!, '\t');
    text += '}\n\n';
    return text;
  }
  text += '\tswitch resp.StatusCode {\n';
  let hasDefault = false;
  for (const kv of foldedMap) {
    if (kv[1].length === 1 && kv[1][0] === 'default') {
      hasDefault = true;
      text += '\tdefault:\n';
    } else {
      text += `\tcase ${formatStatusCodes(kv[1])}:\n`;
    }
    text += generateUnmarshaller(kv[0], '\t\t');
  }
  if (!hasDefault) {
    // add a generic unmarshaller for an unspecified default response
    text += '\tdefault:\n';
    text += generateGenericError();
  }
  text += '\t}\n';
  text += '}\n\n';
  return text;
}

function isArrayOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.DateTime || arrayElem.type === SchemaType.UnixTime;
}

function isArrayOfDate(schema: Schema): boolean {
  if (schema.type !== SchemaType.Array) {
    return false;
  }
  const arraySchema = <ArraySchema>schema;
  const arrayElem = <Schema>arraySchema.elementType;
  return arrayElem.type === SchemaType.Date;
}

function isMapOfDateTime(schema: Schema): boolean {
  if (schema.type !== SchemaType.Dictionary) {
    return false;
  }
  const dictSchema = <DictionarySchema>schema;
  const dictElem = <Schema>dictSchema.elementType;
  return dictElem.type === SchemaType.DateTime || dictElem.type === SchemaType.UnixTime;
}

function isMapOfDate(schema: Schema): boolean {
  if (schema.type !== SchemaType.Dictionary) {
    return false;
  }
  const dictSchema = <DictionarySchema>schema;
  const dictElem = <Schema>dictSchema.elementType;
  return dictElem.type === SchemaType.Date;
}

// returns the media type used by the protocol
function getMediaType(protocol: Protocols): 'JSON' | 'XML' | 'binary' | 'text' | 'form' | 'multipart' | 'none' {
  // TODO: binary, forms etc
  switch (protocol.http!.knownMediaType) {
    case KnownMediaType.Json:
      return 'JSON';
    case KnownMediaType.Xml:
      return 'XML';
    case KnownMediaType.Binary:
      return 'binary';
    case KnownMediaType.Text:
      return 'text';
    case KnownMediaType.Form:
      return 'form';
    case KnownMediaType.Multipart:
      return 'multipart';
    default:
      return 'none';
  }
}

function formatConstantValue(schema: ConstantSchema) {
  // null check must come before any type checks
  if (schema.value.value === null) {
    return 'nil';
  }
  if (schema.valueType.type === SchemaType.String) {
    return `"${schema.value.value}"`;
  }
  return schema.value.value;
}

// returns true if any responses are a binary stream
function hasBinaryResponse(responses: Response[]): boolean {
  for (const resp of values(responses)) {
    if (resp.protocol.http!.knownMediaType === KnownMediaType.Binary) {
      return true;
    }
  }
  return false;
}

// returns the parameters for the public API
// e.g. "ctx context.Context, i int, s string"
function getAPIParametersSig(op: Operation, imports: ImportManager): string {
  const methodParams = getMethodParameters(op);
  const params = new Array<string>();
  if (!isPageableOperation(op) || isLROOperation(op)) {
    imports.add('context');
    params.push('ctx context.Context');
  }
  for (const methodParam of values(methodParams)) {
    params.push(`${(<string>methodParam.language.go!.name).uncapitalize()} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the return signature where each entry is the type name
// e.g. [ '*string', 'error' ]
// apiType describes where the return sig is used.
//   api - for the API definition
//    op - for the operation
// handler - for the response handler
function generateReturnsInfo(op: Operation, apiType: 'api' | 'op' | 'handler'): string[] {
  let returnType = getResponseEnvelopeName(op);
  if (isLROOperation(op)) {
    switch (apiType) {
      case 'handler':
        // we only have a handler for operations that return a schema
        if (isPageableOperation(op)) {
          // we need to consult the final response type name
          returnType = getFinalResponseEnvelopeName(op);
        } else {
          throw new Error(`handler being generated for non-pageable LRO ${op.language.go!.name} which is unexpected`);
        }
        break;
      case 'op':
        returnType = '*http.Response';
        break;
    }
  } else if (isPageableOperation(op)) {
    switch (apiType) {
      case 'api':
      case 'op':
        // pager operations don't return an error
        return [`*${(<PagerInfo>op.language.go!.pageableType).name}`];
    }
  }
  return [returnType, 'error'];
}

function generateLROBeginMethod(op: Operation, imports: ImportManager, isARM: boolean): string {
  const info = <OperationNaming>op.language.go!;
  const params = getAPIParametersSig(op, imports);
  const returns = generateReturnsInfo(op, 'api');
  const clientName = op.language.go!.clientName;
  if (isARM) {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime', 'armruntime');
  }
  let text = '';
  if (hasDescription(op.language.go!)) {
    text += `${comment(`Begin${op.language.go!.name} - ${op.language.go!.description}`, "//", undefined, commentLength)}\n`;
  }
  const zeroResp = getZeroReturnValue(op, 'api');
  text += `func (client *${clientName}) Begin${op.language.go!.name}(${params}) (${returns.join(', ')}) {\n`;
  let opName = op.language.go!.name;
  opName = info.protocolNaming.internalMethod;
  text += `\tresp, err := client.${opName}(${getCreateRequestParameters(op)})\n`;
  text += `\tif err != nil {\n`;
  text += `\t\treturn ${zeroResp}, err\n`;
  text += `\t}\n`;
  text += `\tresult := ${getResponseEnvelopeName(op)}{\n`;
  text += '\t\tRawResponse: resp,\n';
  text += '\t}\n';
  if (isARM) {
    // LRO operation might have a special configuration set in x-ms-long-running-operation-options
    // which indicates a specific url to perform the final Get operation on
    let finalState = '';
    if (op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via']) {
      finalState = op.extensions?.['x-ms-long-running-operation-options']?.['final-state-via'];
    }
    text += `\tpt, err := armruntime.NewPoller("${clientName}.${op.language.go!.name}", "${finalState}", resp, ${getClientPipeline(op)}, client.${info.protocolNaming.errorMethod})\n`;
  } else {
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime');
    text += `\tpt, err := runtime.NewPoller("${clientName}.${op.language.go!.name}",resp, ${getClientPipeline(op)}, client.${info.protocolNaming.errorMethod})\n`;
  }
  text += '\tif err != nil {\n';
  text += `\t\treturn ${zeroResp}, err\n`;
  text += '\t}\n';
  text += `\tresult.Poller = ${emitPoller(op)}`;
  text += `\treturn result, nil\n`;
  // closing braces
  text += '}\n\n';
  return text;
}

