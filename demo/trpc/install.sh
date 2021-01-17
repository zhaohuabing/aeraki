BASEDIR=$(dirname "$0")

kubectl create ns trpc
kubectl label namespace trpc istio-injection=enabled --overwrite=true
kubectl apply -f $BASEDIR/trpc-sample.yaml -n trpc
kubectl apply -f $BASEDIR/serviceentry.yaml -n trpc
kubectl apply -f $BASEDIR/destinationrule.yaml -n trpc
kubectl apply -f $BASEDIR/virtualservice-traffic-splitting.yaml -n trpc