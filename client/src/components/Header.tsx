import Logo from "./Logo";

export const Header = ({ authenticated }: { authenticated: boolean }) => {
  return (
    <div className="flex justify-between items-center">
      <Logo />
      <button className="px-3 py-2 font-rubik bg-gradient-to-b from-indigo-600 to-indigo-700 rounded-md text-sm font-[500] cursor-pointer">
        Sign in with Google
      </button>
    </div>
  );
};
