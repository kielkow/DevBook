$('#form-login').on('submit', signin);

function signin(event) {
    event.preventDefault();

    const email = $('#email').val();
    const password = $('#password').val();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email,
            password,
        }
    }).done(function() {
        window.location = '/home';
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to sign in!', 'error');
    });
}
