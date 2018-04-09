( function( $ ) {
$( document ).ready(function() {
$('#cssmenu').prepend('<div id="menu-button">Menu</div>');
	$('#cssmenu #menu-button').on('click', function(){
		var menu = $(this).next('ul');
		if (menu.hasClass('open')) {
			menu.removeClass('open');
		}
		else {
			menu.addClass('open');
		}
	});
	initLogin();
});
} )( jQuery );

function initLogin(){

    //let userId = getCookie("userId");
    if(0>1){
        userLoggedIn = true;
        $("#sign").hide();
        $("#signout").show();
		$("#order").show();
    }else{
        userLoggedIn = false;
        $("#signout").hide();
		$("#sign").show();
		$("#order").hide();
    }
}

function onSignOut() {
    if(userLoggedIn){
        $("#signout").hide();
        $("#sign").show();
        userLoggedIn = false;
        deleteCookie("userId");
    }

}

function check_pass() {
	if (document.getElementById('password1').value == document.getElementById('password2').value) {
		document.getElementById('password2').style.color="black";
		document.getElementById('submit').disabled = false;
	} else {
		document.getElementById('password2').style.color="red";
		document.getElementById('submit').disabled = true;
	}
}
