param (
    [string]$version,
    [string]$sdkRepoRoot
)

$sdkRepoRoot = Resolve-Path $sdkRepoRoot

Write-Host "Running $version, $sdkRepoRoot"

$buildScriptPath = "$sdkRepoRoot\eng\scripts\build.ps1"

(Get-Content -Raw $buildScriptPath) -replace`
    'goExtension = "@autorest/go@.*"',
    "goExtension = `"@autorest/go@$version`"" | `
    Set-Content $buildScriptPath -NoNewline

Invoke-Expression "$buildScriptPath -skipBuild -cleanGenerated -format -tidy -generate resourcemanager"
