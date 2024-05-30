export default eventHandler(async (event) => {
  const unparsedNpsn = getRouterParam(event, "npsn");

  if (unparsedNpsn === "")
    throw createError({
      status: 400,
      message: "NPSN tidak boleh kosong!",
    });

  const npsn = npsnScheme.safeParse(unparsedNpsn);

  if (!npsn.success)
    throw createError({
      status: 400,
      message: "Format NPSN tidak sesuai!",
    });

  const schoolInfo = await getSpecificSchool(npsn.data);

  if (!schoolInfo)
    throw createError({
      status: 404,
      statusMessage: "Sekolah tidak ditemukan!",
    });

  return { data: schoolInfo };
});
