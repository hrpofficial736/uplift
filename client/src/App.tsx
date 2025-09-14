import Logo from "./components/Logo"
import Prompt from "./components/Prompt"

const App = () => {
  return (
    <div className="p-5 h-screen w-screen relative bg-black text-white font-rubik overflow-hidden">
      <Logo />
      <Prompt />
    </div>
  )
}

export default App
