$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.dislike-publication', dislikePublication);

$('#update-publication').on('click', updatePublication);
$('.delete-publication').on('click', deletePublication);

function createPublication(event) {
    event.preventDefault();

    const title = $('#title').val();
    const content = $('#content').val();

    $.ajax({
        url: '/publications',
        method: 'POST',
        data: {
            title,
            content,
        }
    }).done(function() {
        window.location = '/home';
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to create publication!', 'error');
    });
}

function likePublication(event) {
    event.preventDefault();
    
    const clickedElement = $(event.target);
    const publicationId = clickedElement.closest('div').data('publication-id');

    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/publications/${publicationId}/like`,
        method: 'POST'
    }).done(function() {
        const likesCounter = clickedElement.next('span');
        const likesQuantity = parseInt(likesCounter.text());

        likesCounter.text(likesQuantity + 1);

        clickedElement.addClass('dislike-publication');
        clickedElement.addClass('text-danger');
        clickedElement.removeClass('like-publication');
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to like publication!', 'error');
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}

function dislikePublication(event) {
    event.preventDefault();
    
    const clickedElement = $(event.target);
    const publicationId = clickedElement.closest('div').data('publication-id');

    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/publications/${publicationId}/dislike`,
        method: 'POST'
    }).done(function() {
        const likesCounter = clickedElement.next('span');
        const likesQuantity = parseInt(likesCounter.text());

        likesCounter.text(likesQuantity - 1);

        clickedElement.removeClass('dislike-publication');
        clickedElement.removeClass('text-danger');
        clickedElement.addClass('like-publication');
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to dislike publication!', 'error');
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}

function updatePublication() {
    $(this).prop('disabled', true);

    const publicationId = $(this).data('publication-id');

    const title = $('#title').val();
    const content = $('#content').val();

    $.ajax({
        url: `/publications/${publicationId}`,
        method: 'PUT',
        data: {
            title,
            content,
        }
    }).done(function() {
        Swal.fire('Success!', 'Publication updated!', 'success').then(function() {
            window.location = '/home';
        });
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Ops...', 'Error to update publication!', 'error');
    }).always(function() {
        $('#update-publication').prop('disabled', false);
    });
}

function deletePublication(event) {
    event.preventDefault();

    Swal.fire({
        title: "Attention!",
        text: "Are you sure about delete this publication?",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmation) {
        if (!confirmation.value) return;

        const clickedElement = $(event.target);
        const publication = clickedElement.closest('div');
        const publicationId = publication.data('publication-id');
    
        clickedElement.prop('disabled', true);
    
        $.ajax({
            url: `/publications/${publicationId}`,
            method: 'DELETE'
        }).done(function() {
            publication.fadeOut('slow', function() {
                $(this).remove();
            });
        }).fail(function(error) {
            console.log(error);
            Swal.fire('Ops...', 'Error to delete publication!', 'error');
        });
    });
}
