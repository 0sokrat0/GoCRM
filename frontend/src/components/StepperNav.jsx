import { useState } from "react";
import PropTypes from "prop-types";
import { MenuItem, FormControl, InputLabel, Select, Button, Typography, Box, Card, CardContent, Stepper, Step, StepLabel } from "@mui/material";

const services = [
  { id: 1, name: "Стрижка", price: "1000₽" },
  { id: 2, name: "Покраска", price: "1500₽" },
  { id: 3, name: "Маникюр", price: "1200₽" },
];

const steps = ["Выбор мастера", "Дата и время", "Выбор услуги"];

const StepperNav = ({ activeStep }) => {
  return (
    <Box sx={{ width: "100%", padding: "10px 0" }}>
      <Stepper activeStep={activeStep} alternativeLabel>
        {steps.map((label, index) => (
          <Step key={index}>
            <StepLabel>{label}</StepLabel>
          </Step>
        ))}
      </Stepper>
    </Box>
  );
};

const ServiceSelect = ({ onNext, activeStep }) => {
  const [selectedService, setSelectedService] = useState("");

  return (
    <Box sx={{ display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center", height: "100vh" }}>
      <StepperNav activeStep={activeStep} />
      <Card sx={{ maxWidth: 400, width: "100%", boxShadow: 4, p: 3, textAlign: "center" }}>
        <CardContent>
          <Typography variant="h5" gutterBottom>Выберите услугу</Typography>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Выберите услугу</InputLabel>
            <Select value={selectedService} onChange={(e) => setSelectedService(e.target.value)}>
              {services.map((service) => (
                <MenuItem key={service.id} value={service}>{service.name} - {service.price}</MenuItem>
              ))}
            </Select>
          </FormControl>
          <Button sx={{ mt: 3, width: "100%" }} variant="contained" color="primary" onClick={() => onNext({ service: selectedService })} disabled={!selectedService}>
            Завершить
          </Button>
        </CardContent>
      </Card>
    </Box>
  );
};

ServiceSelect.propTypes = {
  onNext: PropTypes.func.isRequired,
  activeStep: PropTypes.number.isRequired,
};

StepperNav.propTypes = {
  activeStep: PropTypes.number.isRequired,
};

export { ServiceSelect, StepperNav };
