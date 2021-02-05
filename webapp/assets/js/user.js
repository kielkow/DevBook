$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function unfollow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: 'POST'
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to unfollow a user', 'error');
        $('#unfollow').prop('disabled', false);
    });
}

function follow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: 'POST'
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to follow a user', 'error');
        $('#follow').prop('disabled', false);
    });
}

function edit(event) {
    event.preventDefault();

    const name = $('#name').val();
    const email = $('#email').val();
    const nick = $('#nick').val();

    $.ajax({
        url: '/edit-user',
        method: 'PUT',
        data: {
            name,
            email,
            nick,
        }
    }).done(function() {
        Swal.fire('Success', 'User updated with success!', 'success').then(function() {
            window.location = '/profile';
        });
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to edit user!', 'error');
    });
}

function updatePassword(event) {
    event.preventDefault();

    const currentPassword = $('#current-password').val();
    const newPassword = $('#new-password').val();
    const confirmationPassword = $('#confirmation-password').val();

    if (newPassword != confirmationPassword) {
        Swal.fire('Ops...', 'The passwords does not match!', 'warning');
        return;
    }

    $.ajax({
        url: '/update-password',
        method: 'POST',
        data: {
            current: currentPassword,
            new: newPassword,
        }
    }).done(function() {
        Swal.fire('Success', 'Password updated with success!', 'success').then(function() {
            window.location = '/profile';
        });
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to update password!', 'error');
    });
}

function deleteUser() {
    Swal.fire({
        title: 'Attention!',
        text: 'Are sure about that?',
        showCancelButton: true,
        cancelButtonText: 'Cancel',
        icon: 'warning'
    }).then(function(confirmation) {
        if (confirmation.value) {
            $.ajax({
                url: '/delete-user',
                method: 'DELETE',
            }).done(function() {
                Swal.fire('Success', 'Account deleted with success!', 'success').then(function() {
                    window.location = '/logout';
                });
            }).fail(function(error) {
                console.log(error);
                Swal.fire('Ops...', 'Error to delete your account!', 'error');
            });
        }
    });
}
