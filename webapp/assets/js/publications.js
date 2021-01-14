$('#new-publication').on('submit', createPublication);
$('.like-publication').on('click', likePublication);

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
    }).fail(function(error) {
        console.log(error);
        alert("Error to like publication");
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}
