export interface ICaptchaInfo{
    captcha_id: string,
    pic_path: string,
    captcha_length: number
}
export interface IUserInfo{
  id: number
  account: string
  head_pic:string
}

export interface IMenu{
  path: string
  title: string
  icon: string
  header: string
  is_header: number
  children?:IMenu
}
export interface ILoginResponse{
  token: string
  expiresAt: number
  menus: IMenu[]
  unique_auth: string[]
  user_info: IUserInfo
  logo: string
  logo_square: string
  version: string
  newOrderAudioLink:string
}
