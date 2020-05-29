$(document).ready(function () {
    $("#auth").click(
        function () {
            $.ajax({
                url: "/admin/login",
                type: "POST",
                data: $('#ajax_form').serialize(),
                success: function () {
                    location.reload();
                },
                error: function () {
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
                success: function () {
                    window.location.href = "/admin";
                },
                error: function () {
                    $('#result_form').html('<p>Incorrect session</p>');
                }
            });
            return false;
        }
    );
});
