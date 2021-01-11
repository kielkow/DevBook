$('#new-publication').on('submit', createPublication);

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
