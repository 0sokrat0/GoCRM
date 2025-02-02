import { useState } from "react";
import MasterSelect from "./components/MasterSelect";
import DateTimePicker from "./components/DateTimePicker";
import ServiceSelect from "./components/ServiceSelect";

const App = () => {
  const [step, setStep] = useState(1);
  const [data, setData] = useState({});

  const handleNext = (newData) => {
    setData((prev) => ({ ...prev, ...newData }));
    setStep((prev) => prev + 1);
  };

  return (
    <div>
      {step === 1 && <MasterSelect onNext={handleNext} />}
      {step === 2 && <DateTimePicker master={data.master} onNext={handleNext} />}
      {step === 3 && <ServiceSelect onNext={handleNext} />}
      {step === 4 && <h2>Запись подтверждена!</h2>}
    </div>
  );
};

export default App;
