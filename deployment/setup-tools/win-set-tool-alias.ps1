$global:coderoot = (Resolve-Path $PWD\..\..).ToString()

docker build -t deploytools:latest .

Function global:deploytools {
    docker run --rm -it `
    -v "//var/run/docker.sock://var/run/docker.sock" `
    -v $env:userprofile\.aws:/root/.aws `
    -v $env:userprofile\.kube:/root/.kube `
    -v ${coderoot}:/intellimall `
    -v ${PWD}:/intellimall/deployment/.current `
    -w /intellimall/deployment/.current deploytools `
    $args
}

echo "---"
echo ""
echo "Usage: deploytools <cmd [parameters]>"
