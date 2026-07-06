export default function App() {
  return (
    <div className="flex flex-col min-h-screen">
      <main className="flex flex-col flex-8  font-sans bg-dots items-center justify-center" >
        <h1 className=" font-semibold text-4xl text-center ">Make your urls shorter with Shorted</h1>
      </main>

      <hr className=" text-emerald-100"/>
      
      <footer className="flex-1 bg-neutral-800">
        <p className="font-mono opacity-45 text-center ">made by Mr.<a href="https://github.com/MrSplatchy" target="__blank" className=" underline">Splatchy</a></p>
      </footer>
    </div>

  );
}
