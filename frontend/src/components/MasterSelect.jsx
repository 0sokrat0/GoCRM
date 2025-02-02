import { useState, useEffect } from "react";
import PropTypes from "prop-types";
import { MenuItem, Avatar, ListItemText, FormControl, InputLabel, Select, Box, Typography, Card, CardContent, Grid, LinearProgress } from "@mui/material";

const masters = [
  { id: 1, name: "Анна Иванова", description: "Топ-стилист", photo: "https://i.pravatar.cc/300?img=10" },
  { id: 2, name: "Дмитрий Петров", description: "Барбер", photo: "https://i.pravatar.cc/300?img=3" },
  { id: 3, name: "Мария Смирнова", description: "Маникюр", photo: "https://i.pravatar.cc/300?img=5" },
];

// Карточка мастера
const MasterCard = ({ master, theme }) => (
  <Card sx={{ 
    display: "flex", 
    alignItems: "center", 
    p: 2, 
    mb: 3, 
    boxShadow: 2, 
    borderRadius: "12px", 
    maxWidth: 360, 
    mx: "auto",
    backgroundColor: theme.sectionBackground,
    color: theme.textColor
  }}>
    <Avatar src={master.photo} alt={master.name} sx={{ width: 64, height: 64, mr: 2 }} />
    <CardContent sx={{ flex: 1, p: 0 }}>
      <Typography variant="h6" fontWeight="bold">{master.name}</Typography>
      <Typography variant="body2" color="text.secondary">{master.description}</Typography>
    </CardContent>
  </Card>
);

MasterCard.propTypes = { 
  master: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired
};

// Основной компонент
const MasterSelect = ({ onNext }) => {
  const [selectedMaster, setSelectedMaster] = useState(null);
  const [progress, setProgress] = useState(25);
  const [theme, setTheme] = useState({
    backgroundColor: "#ffffff",
    textColor: "#000000",
    sectionBackground: "#f4f4f4",
    buttonColor: "#0088cc",
    buttonTextColor: "#ffffff"
  });

  // Функция для установки темы из `tgWebAppThemeParams`
  const applyTheme = (tg) => {
    const params = tg?.themeParams || {};
    const newTheme = {
      backgroundColor: params.bg_color || "#ffffff",
      textColor: params.text_color || "#000000",
      sectionBackground: params.section_bg_color || params.secondary_bg_color || "#f4f4f4",
      buttonColor: params.button_color || "#0088cc",
      buttonTextColor: params.button_text_color || "#ffffff"
    };
    setTheme(newTheme);

    // Применяем фон и цвет хедера
    tg.setBackgroundColor(newTheme.backgroundColor);
    tg.setHeaderColor("bg_color"); // Используем Telegram-тему для хедера

    // Принудительно меняем фон body
    document.body.style.backgroundColor = newTheme.backgroundColor;
  };

  useEffect(() => {
    if (typeof window.Telegram === "undefined") return;
    const tg = window.Telegram.WebApp;
    
    tg.expand();

    applyTheme(tg);

    tg.MainButton.setParams({
      text: "Далее",
      is_visible: false,
      color: theme.buttonColor,
      text_color: theme.buttonTextColor,
    });

    tg.onEvent("theme_changed", () => applyTheme(tg));
  }, []);

  useEffect(() => {
    if (typeof window.Telegram === "undefined") return;
    const tg = window.Telegram.WebApp;

    if (selectedMaster) {
      tg.MainButton.show();
      tg.MainButton.onClick(() => onNext({ master: selectedMaster }));
      setProgress(50);
    } else {
      tg.MainButton.hide();
      setProgress(25);
    }
  }, [selectedMaster, onNext]);

  return (
    <Box sx={{ textAlign: "center", p: 3, backgroundColor: theme.backgroundColor, color: theme.textColor, minHeight: "100vh" }}>
      <LinearProgress 
        variant="determinate" 
        value={progress} 
        sx={{ 
          mb: 2, 
          height: 5, 
          borderRadius: 5, 
          backgroundColor: theme.sectionBackground 
        }} 
      />

      <Typography variant="h5" fontWeight="bold" sx={{ mb: 2 }}>
        Выберите мастера
      </Typography>

      {selectedMaster && <MasterCard master={selectedMaster} theme={theme} />}

      <FormControl fullWidth sx={{ maxWidth: 360, mx: "auto", textAlign: "left" }}>
        <InputLabel sx={{ color: theme.textColor }}>Выберите мастера</InputLabel>
        <Select
          value={selectedMaster?.id || ""}
          onChange={(e) => setSelectedMaster(masters.find(m => m.id === e.target.value))}
          renderValue={(selected) => {
            const master = masters.find(m => m.id === selected);
            return (
              <Grid container alignItems="center">
                <Avatar src={master.photo} alt={master.name} sx={{ width: 40, height: 40, mr: 2 }} />
                <ListItemText primary={master.name} secondary={master.description} />
              </Grid>
            );
          }}
          sx={{
            backgroundColor: theme.sectionBackground,
            color: theme.textColor,
            borderRadius: "12px"
          }}
        >
          {masters.map((master) => (
            <MenuItem key={master.id} value={master.id}>
              <Grid container alignItems="center">
                <Avatar src={master.photo} alt={master.name} sx={{ width: 40, height: 40, mr: 2 }} />
                <ListItemText primary={master.name} secondary={master.description} />
              </Grid>
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
};

MasterSelect.propTypes = { onNext: PropTypes.func.isRequired };

export default MasterSelect;
