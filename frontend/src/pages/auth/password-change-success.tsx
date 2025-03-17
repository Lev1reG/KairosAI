import StatusCard from "@/components/status-card";

const PasswordChangeSuccessPage = () => {
  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <StatusCard
        type="success"
        headerLabel="Password Change Success"
        message="Your password has been successfully changed."
        buttonLabel="Go to login"
        buttonHref="/login"
      />
    </section>
  );
};

export default PasswordChangeSuccessPage;
