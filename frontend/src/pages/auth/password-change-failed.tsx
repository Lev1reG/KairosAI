import StatusCard from "@/components/status-card";

const PasswordChangeFailedPage = () => {
  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <StatusCard
        type="error"
        headerLabel="Password Change Failed"
        message="Your password could not be changed."
        buttonLabel="Try again"
        buttonHref="/password-change"
      />
    </section>
  );
};

export default PasswordChangeFailedPage;
