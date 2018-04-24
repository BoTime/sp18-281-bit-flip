// check_pass()
function check_pass() {
    if (document.getElementById('password1').value == document.getElementById('password2').value) {
        document.getElementById('password2').style.color="black";
        document.getElementById('submit_signup').disabled = false;
    } else {
        document.getElementById('password2').style.color="red";
        document.getElementById('submit_signup').disabled = true;
    }
}

/**
 * Add "password" field to signup form submission
 */
$('#signup_form').submit(function(e) {
    var password = $('#password1').val();
    var input = $('<input>').attr({
        type: 'hidden',
        name: 'password'
    }).val(password);
    $(this).append(input);
    return true;
});
