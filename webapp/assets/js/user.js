$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);

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
