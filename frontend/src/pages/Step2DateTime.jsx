import { useState } from "react";
import { Button, TextField } from "@mui/material";

// eslint-disable-next-line react/prop-types
const Step2DateTime = ({ onNext, onBack }) => {
  const [dateTime, setDateTime] = useState("");

  return (
    <div>
      <h2>Выберите дату и время</h2>
      <TextField
        type="datetime-local"
        fullWidth
        value={dateTime}
        onChange={(e) => setDateTime(e.target.value)}
      />

      <Button sx={{ marginTop: "20px" }} variant="contained" color="primary" onClick={onNext} disabled={!dateTime}>
        Далее
      </Button>
      <Button sx={{ marginTop: "20px", marginLeft: "10px" }} onClick={onBack}>
        Назад
      </Button>
    </div>
  );
};

export default Step2DateTime;
