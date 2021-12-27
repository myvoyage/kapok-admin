goctl rpc proto -src proto/basic/user.proto -dir ../basic/user -style GoZero
goctl rpc proto -src proto/basic/common.proto -dir ../basic/common -style GoZero


goctl rpc proto -src proto/addons/cms.proto -dir ../addons/cms -style GoZero
goctl rpc proto -src proto/addons/scrm.proto -dir ../addons/scrm -style GoZero

addons              业务服务
    cms             内容服务
    mall            商城服务
    scrm            SCRM服务
    
basic               系统服务&基础服务：如SMS
    common          基础服务：如SMS
    store           商家服务：发布、上传
    trade           交易服务：购物车服务、订单服务
    user            会员中心：会员&商家 服务 登陆、注册等