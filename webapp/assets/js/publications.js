$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.dislike-publication', dislikePublication);

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
        alert("Error to create a publication");
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
        alert("Error to like publication");
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
        alert("Error to dislike publication");
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}
