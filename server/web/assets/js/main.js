$(document).ready(function () {
    $("#auth").click(
        function () {
            $.ajax({
                url: "/admin/login",
                type: "POST",
                data: $('#ajax_form').serialize(),
                success: function (response) {
                    location.reload();
                },
                error: function (response) {
                    $('#result_form').html('<p>Incorrect login or password</p>');
                }
            });
            return false;
        }
    );

    $("#logout").click(
        function () {
            $.ajax({
                url: "/admin/logout",
                type: "POST",
                success: function (response) {
                    window.location.href = "/admin";
                },
                error: function (response) {
                    $('#result_form').html('<p>Incorrect session</p>');
                }
            });
            return false;
        }
    );
});
