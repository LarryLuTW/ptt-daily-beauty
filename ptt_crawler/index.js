/**
 * HTTP Cloud Function.
 * This function is exported by index.js, and is executed when
 * you make an HTTP request to the deployed function's endpoint.
 *
 * @param {Object} req Cloud Function request context.
 * @param {Object} res Cloud Function response context.
 */
exports.getDailyBeauties = (req, res) => {
  console.log(req)
  res.send('Hello World!!!')
}

// functions deploy getDailyBeauties --host "0.0.0.0" --trigger-http
