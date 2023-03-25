github地址：https://gitee.com/zhengpanone/shop-admin.git

npm init vite@latest

cd gin-vue3-admin-web && npm run dev

## eslint

npm install eslint -D
npx eslint --init

## git commit hook

npx mrm@2 lint-staged

https://github.com/conventional-changelog/commitlint#getting-started
npm install --save-dev @commitlint/config-conventional @commitlint/cli
echo "module.exports = {extends: ['@commitlint/config-conventional']}" > commitlint.config.js
