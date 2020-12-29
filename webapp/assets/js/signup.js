$('#form-signup').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    const password = $('#password').val();
    const confirmPassword = $('#confirmation-password').val();
    
    if (password != confirmPassword) {
        alert("Passwords not match!");
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
        alert("Success to create user");
    }).fail(function(error) {
        console.log(error);
        alert("Error to create user");
    });
}
