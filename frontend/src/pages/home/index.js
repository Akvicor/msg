import React, {useEffect, useState} from 'react'
import {message} from 'antd'
import './home.css'
import MarkdownIt from 'markdown-it';
import MdEditor, {PluginComponent} from 'react-markdown-editor-lite';
import 'react-markdown-editor-lite/lib/index.css';
import {userBindHomeTipsFind, userBindHomeTipsSave} from "../../api/userBind";

const mdParser = new MarkdownIt(/* Markdown-it options */);

class EditorView extends PluginComponent {
  static pluginName = 'editor-view'
  static align = 'right'
  static defaultConfig = {view: {menu: true, md: false, html: true}}

  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this)
    this.state = {
      view: this.getConfig('view')
    }
  }

  handleClick() {
    this.state.view.md = false
    this.props.editor.setView(this.state.view)
    userBindHomeTipsSave({content: this.props.editor.getMdValue()}).then(({data}) => {
      if (data.code !== 0) {
        message.open({
          type: 'warning',
          content: '保存失败'
        })
        return
      }
      message.open({
        type: 'success',
        content: '保存成功'
      })
    })
  }

  render() {
    return (
      <span
        className="button button-type-update"
        title="Update"
        onClick={this.handleClick}
      >
        更新
      </span>
    );
  }
}

MdEditor.use(EditorView)

const Home = () => {
  const [editorValue, setEditorValue] = useState('## Hi')

  useEffect(() => {
    userBindHomeTipsFind().then(({data}) => {
      if (data.code !== 0) {
        message.open({
          type: 'warning',
          content: '获取Tips失败'
        })
        return
      }
      setEditorValue(data.data.content)
    })
  }, []);

  const handleEditorChange = ({html, text}) => {
    setEditorValue(text)
  }

  return (
    <MdEditor value={editorValue}
              view={EditorView.defaultConfig.view}
              canView={{menu: true, md: true, html: true, both: true, fullScreen: false, hideMenu: true}}
              renderHTML={text => mdParser.render(text)} onChange={handleEditorChange}/>
  )
}

export default Home
