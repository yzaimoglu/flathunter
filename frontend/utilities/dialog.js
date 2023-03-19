import Swal from 'sweetalert2';

export const Dialog = async (steps) => {
    const swalQueueStep = Swal.mixin({
        confirmButtonText: 'Weiter',
        confirmButtonColor: '#3B82F6',
        cancelButtonText: 'Zurück',
        progressSteps: steps,
        input: 'text',
        inputAttributes: {
            required: true
        },
        reverseButtons: true,
        validationMessage: 'Dieses Feld ist Pflicht'
    })

    const values = []
    let currentStep

    for (currentStep = 0; currentStep < steps.length;) {
        const confirmButtonText = currentStep === steps.length - 1 ? 'Beenden' : 'Weiter'
        const result = await swalQueueStep.fire({
            title: `Frage ${steps[currentStep]}`,
            inputValue: values[currentStep],
            showCancelButton: currentStep > 0,
            confirmButtonText: confirmButtonText,
            currentProgressStep: currentStep
        })

        if (result.value) {
            values[currentStep] = result.value
            currentStep++
        } else if (result.dismiss === Swal.DismissReason.cancel) {
            currentStep--
        } else {
            break
        }
    }

    if (currentStep === steps.length) {
        return values
    }
}