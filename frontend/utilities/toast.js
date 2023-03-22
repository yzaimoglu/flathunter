import Swal from 'sweetalert2';

export const ToastTR = Swal.mixin({
    toast: true,
    position: 'top-right',
    iconColor: 'white',
    customClass: {
        popup: 'colored-toast'
    },
    showConfirmButton: false,
    timer: 1500,
    timerProgressBar: true
});

export const ToastTL = Swal.mixin({
    toast: true,
    position: 'top-left',
    iconColor: 'white',
    customClass: {
        popup: 'colored-toast'
    },
    showConfirmButton: false,
    timer: 1500,
    timerProgressBar: true
});

export const ToastBR = Swal.mixin({
    toast: true,
    position: 'bottom-right',
    iconColor: 'white',
    customClass: {
        popup: 'colored-toast'
    },
    showConfirmButton: false,
    timer: 1500,
    timerProgressBar: true
});

export const ToastBL = Swal.mixin({
    toast: true,
    position: 'bottom-left',
    iconColor: 'white',
    customClass: {
        popup: 'colored-toast'
    },
    showConfirmButton: false,
    timer: 1500,
    timerProgressBar: true
});

export const Toast = async (position, title, icon) => {
    switch (position) {
        case 'tr':
            await ToastTR.fire({
                title: title,
                icon: icon
            });
            break;
        case 'tl':
            await ToastTL.fire({
                title: title,
                icon: icon
            });
            break;
        case 'br':
            await ToastBR.fire({
                title: title,
                icon: icon
            });
            break;
        case 'bl':
            await ToastBL.fire({
                title: title,
                icon: icon
            });
            break;
        default:
            await ToastTR.fire({
                title: title,
                icon: icon
            });
            break;
        }
}
