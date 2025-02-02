import { useState, useEffect } from "react";
import PropTypes from "prop-types";
import { MenuItem, FormControl, InputLabel, Select, Typography, Box } from "@mui/material";

const services = [
  { id: 1, name: "Стрижка", price: "1000₽" },
  { id: 2, name: "Покраска", price: "1500₽" },
  { id: 3, name: "Маникюр", price: "1200₽" },
];

const ServiceSelect = ({ onNext }) => {
  const [selectedService, setSelectedService] = useState("");

  useEffect(() => {
    const tg = window.Telegram.WebApp;
    tg.MainButton.setParams({
      text: "Завершить",
      is_visible: false,
    });
  }, []);

  useEffect(() => {
    const tg = window.Telegram.WebApp;
    if (selectedService) {
      tg.MainButton.show();
      tg.MainButton.onClick(() => onNext({ service: selectedService }));
    } else {
      tg.MainButton.hide();
    }
  }, [selectedService, onNext]);

  return (
    <Box sx={{ textAlign: "center", p: 3 }}>
      <Typography variant="h5" fontWeight="bold" sx={{ mb: 2 }}>
        Выберите услугу
      </Typography>
      <FormControl fullWidth>
        <InputLabel>Выберите услугу</InputLabel>
        <Select value={selectedService} onChange={(e) => setSelectedService(e.target.value)}>
          {services.map((service) => (
            <MenuItem key={service.id} value={service}>
              {service.name} - {service.price}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
};

ServiceSelect.propTypes = { onNext: PropTypes.func.isRequired };

export default ServiceSelect;
