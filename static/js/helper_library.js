function showLoader () {

    $('.loader').show();

}

function hideLoader () {

    $('.loader').hide();

}

function showErrorMessage (title, message) {

    $('.alert-danger').remove();

    var alertMsg = '<div class="alert alert-danger alert-dismissible" role="alert">';

    alertMsg += '<button type="button" class="close" data-dismiss="alert" aria-label="Close">';
    alertMsg += '<span aria-hidden="true">&times;</span>';
    alertMsg += '</button>';
    alertMsg += '<h4 class="alert-heading">' + title + '</h4>';
    alertMsg += '<p class="message">' + message + '</p>';
    alertMsg += '</div>';


    $('.layout').find('.container').prepend(alertMsg);
    
    

    $('.alert-danger').alert();

    setTimeout(function () {
        $('.alert-danger').remove();
    }, 10 * 1000);

}

function toTitleCase(str) {
    return str.replace(/\w\S*/g, function(txt){return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();});
}