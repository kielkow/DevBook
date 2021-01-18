$('#form-signup').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    const password = $('#password').val();
    const confirmPassword = $('#confirmation-password').val();
    
    if (password != confirmPassword) {
        Swal.fire('Ops...', 'Passwords does not match!', 'error');
        return;
    }

    const name = $('#name').val();
    const email = $('#email').val();
    const nick = $('#nick').val();

    $.ajax({
        url: '/users',
        method: 'POST',
        data: {
            name,
            email,
            nick,
            password,
        }
    }).done(function() {
        Swal.fire('Success!', 'User created!', 'success').then(function() {
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
        });
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to create user!', 'error');
    });
}
