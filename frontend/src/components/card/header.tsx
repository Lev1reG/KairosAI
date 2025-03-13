interface HeaderProps {
  label: string;
}

const Header = ({ label }: HeaderProps) => {
  return (
    <div className="w-full flex justify-center items-center">
      <h2 className="font-semibold text-2xl text-black">{label}</h2>
    </div>
  );
};

export default Header;
