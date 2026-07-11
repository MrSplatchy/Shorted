import { useState } from "react";
import axios from "axios";

export default function App() {
  let tab= [
    { label: "Create Shortcut", content: <CreateURL key="key"/> },
    // { label: "Change URL Shortcut", content: <p>Second tab</p> },
  ];
  return (
    <div className="flex flex-col min-h-screen">
      <main className="flex flex-col flex-9 h-full  font-sans bg-dots items-center justify-evenly " >
        <h1 className="font-semibold text-6xl text-center ">Shorted</h1>

        <TabbedCard tabs={tab}/>
      </main>

      <hr className=" text-emerald-100"/>
      
      <footer className="flex-1 bg-neutral-800">
        <p className="font-mono opacity-45 text-center ">Made by Mr.<a href="https://github.com/MrSplatchy" target="__blank" className=" underline">Splatchy</a></p>
      </footer>
    </div>

  );
}

// Returns a map of buttons and their content
function TabbedCard({ tabs }) {
  const [activeIndex, setActiveIndex] = useState(0)

  const activeContent = tabs[activeIndex]?.content

  return (
    <div className=" w-11/12 max-w-2xl overflow-hidden rounded-l border border-border bg-stone-800 shadow-xl shadow-black/20 ">
      <div className="flex border-b border-border">
        {tabs.map((tab, index) => (
          <button
            key={index}
            onClick={() => setActiveIndex(index)}
            className={`relative px-4 py-2 text-sm font-medium w-screen transition-colors border-b-2 ${
              activeIndex === index
                ? "text-foreground border-green-600 border-b-4 "
                : "text-muted-foreground hover:text-foreground border-transparent hover:cursor-crosshair hover:bg-green-600"
            }`}
          >
            {tab.label}
          </button>
        ))}
      </div>

      <div className="p-6 text-card-foreground ">{activeContent}</div>
    </div>
  )
}

function CreateURL(){
  const [url, setURL] = useState("")
  const [submitted, setSubmitted] = useState(false)
  const [state, setState] = useState("")

  const apiUrl = import.meta.env.VITE_API_URL || "http://localhost:8080"
  const redirectUrl = apiUrl+"/"+state

  const handleChange = event => setURL(event.target.value)

  const handleSubmit = event => {
    event.preventDefault()

  axios
    .post(`${apiUrl}/api/`, {
      url: url
    })
    .then((response) => setState(response.data.shortCode))
    .catch((error) => setState(error.message))
    .finally(() => console.log("Request completed"));

    setSubmitted(true)

    return state
  }

  return(
    <form className="flex flex-col  items-center justify-center min-w-7/12" onSubmit={handleSubmit}>
      <input
        type="text" 
        value={url}
        onChange={handleChange}
        placeholder="Please enter an URL to shorten"
        className="min-w-7/12 placeholder:text-center text-center"
      />

      <button type="submit" className="border w-5/12 mt-4 field-sizing-content">
          Submit
      </button>

      {submitted && (
        <a className=" mt-3 " href={redirectUrl}>{state}</a>
      )}
    </form>
  )
}