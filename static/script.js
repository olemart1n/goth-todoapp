window.addEventListener("DOMContentLoaded", () => {
  const headerLinks = document.querySelectorAll("header a");
  const currentPath = window.location.pathname;
  const pageName = currentPath.split("/").pop();
  console.log(headerLinks);
  headerLinks.forEach((link) => {
    console.log(
      link.innerHTML.trim().toLowerCase() +
        " -> " +
        pageName.trim().toLowerCase(),
    );
    if (
      link.innerHTML.trim().toLowerCase().replace(/\s/g, "") ===
      pageName.trim().toLowerCase().replace(/\s/g, "")
    ) {
      link.style.textDecoration = "underline";
    }
    console.log(link);
  });
});
