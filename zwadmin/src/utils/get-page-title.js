import defaultSettings from '@/settings'

const title = defaultSettings.title || '磨刀4'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
