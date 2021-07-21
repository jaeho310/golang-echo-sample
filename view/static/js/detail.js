$(document).ready(function(){
    console.log("detail ready");
    getUser($("#template_id").html())
});

let user = new Object()

function getUser(id) {
    $.ajax({
        url: '/api/users/' + id,
        contentType: 'application/json',
        type: 'get',
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            user = data
            $("#id").val(data.id)
            $("#name").val(data.name)
            $("#create").val(data.createdAt)
            $("#update").val(data.updatedAt)
        },
        error: function(request,status,error){
            alert(" message = " + request.responseText);
        }
    })
}

$("#delete_btn").on("click", function() {
    let result = confirm("삭제하시겠습니까?");
    if (result == false) {
        return;
    }
    let id = $("#id").val();
    $.ajax({
        url: '/api/users/'+ id,
        contentType: 'application/json',
        type: 'delete',
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            alert("삭제성공");
            window.location.href="/list"
        },
        error: function(request,status,error){
            alert(" message = " + request.responseText);
        }
    })
});

$("#cancel_btn").on("click", function() {
    window.location.href="/list"
});

$("#update_btn").on("click", function() {
    if (user["name"] == $("#name").val()) {
        alert("변경사항이 없습니다.")
        return;
    }
    let result = confirm("변경하시겠습니까?");
    if (result == false) {
        return;
    }

    user["name"] = $("#name").val();

    $.ajax({
        url: '/api/users',
        contentType: 'application/json',
        type: 'patch',
        data: JSON.stringify(user),
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            alert("변경완료");
            window.location.reload(true)
        },
        error: function(request,status,error){
            alert(" message = " + request.responseText);
        }
    })
});