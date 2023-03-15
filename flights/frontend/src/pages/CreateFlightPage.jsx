import { Button, Container, Grid } from "@mui/material";
import CreateFlightForm from "components/flight_management/CreateFlightForm";
import { Form } from "react-final-form";
import REGEX from "regex";

const numberRegex = new RegExp(REGEX.NUMBER)
const CreateFlightPage = () => {

    const validate = (values) => {
        let returnObject = {}
        if (!values.departure) {
            returnObject.departure = 'This field is required!'
        }
        if (!values.destination) {
            returnObject.destination = 'This field is required!'
        }
        if (new Date(values.departureDateTime) > new Date(values.arrivalDateTime)) {
            returnObject.departure = 'Arrival must be after departure!'
            returnObject.destination = 'Arrival must be after departure!'
        }
        if (!numberRegex.test(values.price)) {
            returnObject.price = 'This must be a number!'
        }
        if (!numberRegex.test(values.seats)) {
            returnObject.seats = 'This must be a number!'
        }
        return returnObject
    }

    const initialValues = {
        departureDateTime: new Date(),
        arrivalDateTime: new Date()
    }

    const onSubmit = (values) => {
        //TODO: make request to API for flight creation
        console.log(values);
    }

    return (
        <>
            <Form
                onSubmit={onSubmit}
                validate={validate}
                initialValues={initialValues}
                render={({ handleSubmit, values }) => (
                    <form onSubmit={handleSubmit} noValidate>
                        <Grid container sx={{ margin: 'auto' }}>
                            <Grid item xs={12}>
                                <Container sx={{ display: 'grid', placeItems: 'center', width: '90%', marginTop: '4%' }}>
                                    <CreateFlightForm />
                                    <Button variant="contained" color="primary" type='submit' sx={{marginTop: '1rem'}}>
                                        Submit
                                    </Button>
                                </Container>
                            </Grid>
                        </Grid>
                    </form>)}
            >
            </Form>
        </>
    );
}

export default CreateFlightPage;