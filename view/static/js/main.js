// let token = $("meta[name='_csrf']").attr("content");
// let header = $("meta[name='_csrf_header']").attr("content");
let fileList = new Object();

$(document).ready(function(){
    console.log("main ready");

});

$("#reg_btn").on("click", function() {
    // console.log("test");
    let name = $("#userName").val();
    console.log(name);
    var data = new Object();
    data["name"] = name;
    $.ajax({
        url: '/api/users',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(data),
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            alert("등록성공");
            $("#userName").val("");
        },
        error: function(request,status,error){
            alert(" message = " + request.responseText);
        }
    })

});
