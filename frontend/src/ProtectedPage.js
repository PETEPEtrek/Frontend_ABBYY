import Sidebar from "./components/Sidebar"
import MainContent from "./components/MainContent"
import TagsPanel from "./components/TagsPanel"
const ProtectedPage = () => {
  return (
    <div className="flex">
            <Sidebar/>
            <MainContent/>
            <TagsPanel/>
    </div>
  )
}

export default ProtectedPage 