# data-distribute
beego，gopher-lua(https://github.com/yuin/gopher-lua)

- 业务平台将产生的数据通过接口发送至分发平台；
- 分发平台配置数据的订阅方，并调用订阅方提供的接口发送数据；
- 每家订阅方的接口都不一样，有http的，有webservice的，各是各的定义和实现，所以分发平台使用lua脚本去调用订阅方的接口，以实现定制化。
