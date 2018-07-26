$(function(){

})
function login() {
    var form_data = $("#loginForm").serialize();
    console.info(form_data);
    $.ajax({
        url:"http://127.0.0.1:9090/login",
        data:form_data,
        method:"post",
        async:true,
        success:function (data) {
            console.info(data);
            var json = JSON.parse(data);
            if(json.success=="ok"){
                alert("登录成功");
            }else{
                alert("登录失败");
            }
        }
    });
}