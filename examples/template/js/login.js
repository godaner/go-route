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
            alert(json.msg);
            if(json.code==1){
                location.href="http://127.0.0.1:9090/html/home.html"
            }
        }
    });
}