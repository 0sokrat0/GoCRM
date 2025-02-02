/* eslint-disable react/prop-types */
import MasterSelect from "../components/MasterSelect";
import DateTimePicker from "../components/DateTimePicker";
import ServiceSelect from "../components/ServiceSelect";

// eslint-disable-next-line react/prop-types
const BookingPage = ({ step, onNext, bookingData }) => {
  return (
    <div style={{display: 'flex', justifyContent: 'center'}}>
      {step === 0 && <MasterSelect onNext={onNext} />}
      {step === 1 && <DateTimePicker master={bookingData.master} onNext={onNext} />}
      {step === 2 && <ServiceSelect onNext={onNext} />}
      {step === 3 && <h2>Запись подтверждена!</h2>}
    </div>
  );
};

export default BookingPage;
