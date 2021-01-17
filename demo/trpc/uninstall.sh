BASEDIR=$(dirname "$0")

kubectl delete -f $BASEDIR/trpc-sample.yaml -n trpc
kubectl delete -f $BASEDIR/serviceentry.yaml -n trpc
kubectl delete -f $BASEDIR/destinationrule.yaml -n trpc
kubectl delete -f $BASEDIR/virtualservice-traffic-splitting.yaml -n trpc
kubectl delete ns trpc