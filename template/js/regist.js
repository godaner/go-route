$(function(){

})
function regist() {
    var form_data = $("#registForm").serialize();
    console.info(form_data);
    $.ajax({
        url:"http://127.0.0.1:9090/regist",
        data:form_data,
        method:"post",
        async:true,
        success:function (data) {
            console.info(data);
            var json = JSON.parse(data);
            alert(json.msg);
        }
    });
}