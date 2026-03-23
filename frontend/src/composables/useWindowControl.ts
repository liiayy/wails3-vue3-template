import { Window } from '@wailsio/runtime'

/**
 * useWindowControl — 当前窗口控制工具箱
 *
 * 封装 Wails 3 Window API，提供最小化、最大化、关闭、全屏等操作。
 * 无需手动管理窗口实例引用。
 *
 * @example
 * ```ts
 * const { minimise, toggleMaximise, close } = useWindowControl()
 *
 * // 在模板中
 * <button @click="minimise">—</button>
 * <button @click="toggleMaximise">□</button>
 * <button @click="close">×</button>
 * ```
 */
export function useWindowControl() {
  const win = Window

  async function minimise() {
    await win.Minimise()
  }

  async function maximise() {
    await win.Maximise()
  }

  async function unmaximise() {
    await win.UnMaximise()
  }

  async function toggleMaximise() {
    await win.ToggleMaximise()
  }

  async function close() {
    await win.Close()
  }

  async function fullscreen() {
    await win.Fullscreen()
  }

  async function toggleFullscreen() {
    await win.ToggleFullscreen()
  }

  async function center() {
    await win.Center()
  }

  async function setTitle(title: string) {
    await win.SetTitle(title)
  }

  async function setSize(width: number, height: number) {
    await win.SetSize(width, height)
  }

  async function isMaximised(): Promise<boolean> {
    return await win.IsMaximised()
  }

  async function isFullscreen(): Promise<boolean> {
    return await win.IsFullscreen()
  }

  return {
    minimise,
    maximise,
    unmaximise,
    toggleMaximise,
    close,
    fullscreen,
    toggleFullscreen,
    center,
    setTitle,
    setSize,
    isMaximised,
    isFullscreen,
  }
}
