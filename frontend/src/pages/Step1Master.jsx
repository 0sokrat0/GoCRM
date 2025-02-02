import  { useState } from "react";
import { MenuItem, Avatar, ListItemText, ListItemAvatar, FormControl, InputLabel, Select, Button } from "@mui/material";

const masters = [
  { id: 1, name: "Анна Иванова", description: "Топ-стилист", photo: "https://i.pravatar.cc/300?img=10" },
  { id: 2, name: "Дмитрий Петров", description: "Барбер", photo: "https://i.pravatar.cc/300?img=3" },
  { id: 3, name: "Мария Смирнова", description: "Маникюр", photo: "https://i.pravatar.cc/300?img=5" },
];

// eslint-disable-next-line react/prop-types
const Step1Master = ({ onNext }) => {
  const [selectedMaster, setSelectedMaster] = useState("");

  return (
    <div>
      <h2>Выберите мастера</h2>
      <FormControl fullWidth>
        <InputLabel>Выберите мастера</InputLabel>
        <Select value={selectedMaster} onChange={(e) => setSelectedMaster(e.target.value)}>
          {masters.map((master) => (
            <MenuItem key={master.id} value={master.id}>
              <ListItemAvatar>
                <Avatar src={master.photo} alt={master.name} />
              </ListItemAvatar>
              <ListItemText primary={master.name} secondary={master.description} />
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <Button sx={{ marginTop: "20px" }} variant="contained" color="primary" onClick={onNext} disabled={!selectedMaster}>
        Далее
      </Button>
    </div>
  );
};

export default Step1Master;
