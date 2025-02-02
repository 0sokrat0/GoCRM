import { useState } from "react";
import PropTypes from "prop-types";
import { MenuItem, FormControl, InputLabel, Select, Typography, Box } from "@mui/material";
import { LocalizationProvider, DatePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

const availableTimes = {
  1: ["10:00", "12:00", "14:00"],
  2: ["11:00", "15:00", "18:00"],
  3: ["09:00", "13:00", "17:00"],
};

const DateTimePicker = ({ master, onNext }) => {
  const [selectedDate, setSelectedDate] = useState(null);
  const [selectedTime, setSelectedTime] = useState("");

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Box sx={{ textAlign: "center", p: 3 }}>
        <Typography variant="h5" fontWeight="bold" sx={{ mb: 2 }}>
          Выберите дату и время
        </Typography>
        <DatePicker
          label="Выберите дату"
          value={selectedDate}
          onChange={setSelectedDate}
          sx={{ width: "100%", mb: 2 }}
        />
        <FormControl fullWidth>
          <InputLabel>Выберите время</InputLabel>
          <Select value={selectedTime} onChange={(e) => setSelectedTime(e.target.value)}>
            {(availableTimes[master?.id] || []).map((time) => (
              <MenuItem key={time} value={time}>{time}</MenuItem>
            ))}
          </Select>
        </FormControl>
        <button
          style={{
            marginTop: "20px",
            padding: "10px 20px",
            borderRadius: "8px",
            backgroundColor: "#007bff",
            color: "#fff",
            border: "none",
            cursor: "pointer",
            fontSize: "16px",
          }}
          onClick={() => onNext({ dateTime: { date: selectedDate, time: selectedTime } })}
          disabled={!selectedDate || !selectedTime}
        >
          Далее
        </button>
      </Box>
    </LocalizationProvider>
  );
};

DateTimePicker.propTypes = { master: PropTypes.object, onNext: PropTypes.func.isRequired };

export default DateTimePicker;
