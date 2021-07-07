```shell
# docker核心配置
# cat /etc/docker/daemon.json
{
    "exec-opts": ["native.cgroupdriver=systemd"],
    "log-driver": "json-file",
    "log-opts": {
        "max-size": "100m",
        "max-file": "10"
    },
    "oom-score-adjust": -1000,
    "registry-mirrors": ["https://pqbap4ya.mirror.aliyuncs.com"],
    "storage-driver": "overlay2",
    "max-concurrent-downloads": 10,
    "max-concurrent-uploads": 5,
    "storage-opts":["overlay2.override_kernel_check=true"],
    "live-restore": true
}
# 查看pods
# kubectl get pods -o wide
NAME                              READY   STATUS    RESTARTS   AGE     IP               NODE            NOMINATED NODE   READINESS GATES
deployment-oss-7b4fcfdd66-6n5qb   1/1     Running   0          5m21s   10.244.33.177    172-16-100-62   <none>           <none>
# 获取请求日志
# kubectl logs -f deployment-oss-7b4fcfdd66-6n5qb
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2021/07/07 00:06:24 [notice] 1#1: using the "epoll" event method
2021/07/07 00:06:24 [notice] 1#1: nginx/1.21.0
2021/07/07 00:06:24 [notice] 1#1: built by gcc 8.3.0 (Debian 8.3.0-6) 
2021/07/07 00:06:24 [notice] 1#1: OS: Linux 4.19.0-13-amd64
2021/07/07 00:06:24 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2021/07/07 00:06:24 [notice] 1#1: start worker processes
2021/07/07 00:06:24 [notice] 1#1: start worker process 30
2021/07/07 00:06:24 [notice] 1#1: start worker process 31
10.244.69.64 - - [07/Jul/2021:00:07:03 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.64.0" "-"
10.244.69.64 - - [07/Jul/2021:00:07:33 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.64.0" "-"
# 使用docker获取实例ID
# docker ps | grep oss
d0d60bf14732   nginx                                                                 "/docker-entrypoint.…"   6 minutes ago   Up 6 minutes             k8s_nginx_deployment-oss-7b4fcfdd66-6n5qb_default_a9240d5a-f262-4021-98c8-08136bf892e2_0
# 根据实例ID获取POD临时日志文件
# tail -f /var/lib/docker/containers/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa-json.log 
{"log":"2021/07/07 00:06:24 [notice] 1#1: using the \"epoll\" event method\n","stream":"stderr","time":"2021-07-07T00:06:24.784338402Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: nginx/1.21.0\n","stream":"stderr","time":"2021-07-07T00:06:24.784519289Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: built by gcc 8.3.0 (Debian 8.3.0-6) \n","stream":"stderr","time":"2021-07-07T00:06:24.784545378Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: OS: Linux 4.19.0-13-amd64\n","stream":"stderr","time":"2021-07-07T00:06:24.784564094Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576\n","stream":"stderr","time":"2021-07-07T00:06:24.784600703Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: start worker processes\n","stream":"stderr","time":"2021-07-07T00:06:24.78469815Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: start worker process 30\n","stream":"stderr","time":"2021-07-07T00:06:24.784870668Z"}
{"log":"2021/07/07 00:06:24 [notice] 1#1: start worker process 31\n","stream":"stderr","time":"2021-07-07T00:06:24.785004889Z"}
{"log":"10.244.69.64 - - [07/Jul/2021:00:07:03 +0000] \"GET / HTTP/1.1\" 200 612 \"-\" \"curl/7.64.0\" \"-\"\n","stream":"stdout","time":"2021-07-07T00:07:03.663354683Z"}
{"log":"10.244.69.64 - - [07/Jul/2021:00:07:33 +0000] \"GET / HTTP/1.1\" 200 612 \"-\" \"curl/7.64.0\" \"-\"\n","stream":"stdout","time":"2021-07-07T00:07:33.257250118Z"}
# 统计临时文件占用磁盘大小
# du -sh /var/lib/docker/containers/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa-json.log
4.0K    /var/lib/docker/containers/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa-json.log
# 根据kubelet日志，实例ID及pods
I0707 08:05:52.426895   22571 kubelet.go:1868] SyncLoop (UPDATE, "api"): "deployment-oss-7b4fcfdd66-6n5qb_default(a9240d5a-f262-4021-98c8-08136bf8
92e2)"
I0707 08:05:52.496130   22571 provider.go:102] Refreshing cache for provider: *credentialprovider.defaultDockerConfigProvider
I0707 08:06:19.631072   22571 kube_docker_client.go:344] Pulling image "nginx:1.21.0": "b4d181a07f80: Downloading [===============================
=====>              ]  19.56MB/27.15MB"
I0707 08:06:24.384641   22571 kube_docker_client.go:347] Stop pulling image "nginx:1.21.0": "Status: Downloaded newer image for nginx:1.21.0"
I0707 08:06:25.759647   22571 kubelet.go:1899] SyncLoop (PLEG): "deployment-oss-7b4fcfdd66-6n5qb_default(a9240d5a-f262-4021-98c8-08136bf892e2)", e
vent: &pleg.PodLifecycleEvent{ID:"a9240d5a-f262-4021-98c8-08136bf892e2", Type:"ContainerStarted", Data:"d0d60bf147325e56b0e98707f12123bc52250fd637
77ca439205ef918d87b5aa"}
# du -sh /var/lib/kubelet/pods/a9240d5a-f262-4021-98c8-08136bf892e2/
21K     /var/lib/kubelet/pods/a9240d5a-f262-4021-98c8-08136bf892e2/

# 删除PODS
# kubectl delete -f deploy.yaml
deployment.apps "deployment-oss" deleted
# 目录被删除
# du -sh /var/lib/kubelet/pods/a9240d5a-f262-4021-98c8-08136bf892e2/
du: cannot access '/var/lib/kubelet/pods/a9240d5a-f262-4021-98c8-08136bf892e2/': No such file or directory
# 临时文件被删除
# ls /var/lib/docker/containers/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa-json.log  
ls: cannot access '/var/lib/docker/containers/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa/d0d60bf147325e56b0e98707f12123bc52250fd63777ca439205ef918d87b5aa-json.log': No such file or director
```