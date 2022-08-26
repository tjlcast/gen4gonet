# 根据service的struct生成rpc的服务端
./gen4gonet -f {srvFile} -g rpcs
# 根据service的struct生成rpc的客户端
./gen4gonet -f {srvFile} -g rpcc
#  根据service的struct生成rest的服务端
./gen4gonet -f {srvFile} -g rests