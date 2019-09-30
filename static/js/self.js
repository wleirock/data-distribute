/**
 * 点击侧栏条目时，动态修改其样式
 */
function changeSideStyle(){
    var identity = document.getElementById('identity');
    var dd = document.querySelector('#'+identity.value);
    dd.className = 'layui-this';
    dd.parentNode.parentNode.className += ' layui-nav-itemed';
}

/**
 * 正则判断文件名称是否是符合要求的lua文件名
 * @param {文件名称} fileName 
 */
function ifLuafile(fileName){
    var regu = /[a-zA-Z0-9_]{2,20}.lua$/;
    var re = new RegExp(regu);
    if(re.test(fileName)){
        return true;
    }
    return false;
}