const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const {
    IgnorePlugin,
    DefinePlugin,
} = require('webpack');
const extractStyles = require('./extractStyles');

const MAIN_SRC_PATH = './src';

/**
 * @param options {Object}
 * @param options.isProduction {Boolean}
 * @param options.buildFolder {String}
 * @param options.appVersion {String}
 * @param options.extractStylesFile {Boolean}
 */
module.exports = (options) => {
    return {
        mode: options.isProduction ? 'production' : 'development',
        entry: {
            bundle: `${MAIN_SRC_PATH}/index.js`,
        },
        output: {
            path: `${process.cwd()}/${options.buildFolder}`,

            // @docs https://webpack.js.org/guides/caching/#deterministic-hashes
            filename: options.isProduction ?
                './[name]-[chunkhash].js' :
                './[name].js',
            chunkFilename: options.isProduction ?
                './[id].chunk-[chunkhash].js' :
                './[id].chunk.js',
            publicPath: '',
        },
        resolve: {
            extensions: ['.js', '.ts'],
        },
        module: {
            rules: [
                {
                    test: /\.(t|j)sx?$/,
                    exclude: /node_modules/,
                    use: 'ts-loader',
                },

                extractStyles.moduleRule(options.extractStylesFile),
            ],
        },
        plugins: [
            // Ignoring warnings for following plugins, case they doesn't matter
            new IgnorePlugin(/regenerator|nodent|js-beautify/, /ajv/),

            // Defining global ENV variable
            new DefinePlugin({
                ENV: {production: options.isProduction},
            }),

            new HtmlWebpackPlugin({
                template: `${MAIN_SRC_PATH}/index.ejs`,
                filename: './index.html',
                appVersion: options.appVersion,
                appBuildDate: new Date(),
            }),

            new CleanWebpackPlugin({
                verbose: true,
                dry: false,
                cleanOnceBeforeBuildPatterns: [
                    '**/*',
                    '!.gitignore',
                ],
            }),

            new CopyPlugin([
                `${MAIN_SRC_PATH}/wasm_exec.js`
            ]),

            ...extractStyles.plugins(options.extractStylesFile, options.isProduction)
        ],
    };
}
