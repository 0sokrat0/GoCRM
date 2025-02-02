
import { Button } from "@mui/material";

// eslint-disable-next-line react/prop-types
const Step3Service = ({ onBack }) => {
  return (
    <div>
      <h2>Выбор услуги</h2>
      <p>Тут будет список доступных услуг.</p>

      <Button sx={{ marginTop: "20px" }} onClick={onBack}>
        Назад
      </Button>
    </div>
  );
};

export default Step3Service;
