// @flow
import * as React from 'react'
import type {Props} from './oriented-image.types'
import {Image} from '../common-adapters'

// OrientedImage will render a NativeImage.
// However, Oriented image is never rendered from a native component as mobile
// devices accurately read and respect EXIF Orientation data.
// This is done because the parent(s) of OrientedImage are split.
export default ({src, style}: Props) => {
  return <Image src={src} style={style} />
}
