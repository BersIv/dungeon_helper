package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate1Binding
import com.example.dungeon_helper.shared.JsonHelper
import com.example.dungeon_helper.shared.SharedViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.OkHttpClient
import okhttp3.Request
import java.io.IOException

class CharacterCreate1 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate1()
    }

    private lateinit var viewModel: CharacterCreate1ViewModel
    private var _binding: FragmentCharacterCreate1Binding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate1ViewModel = ViewModelProvider(this)[CharacterCreate1ViewModel::class.java]
        _binding = FragmentCharacterCreate1Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView: TextView = binding.textCharacterCreate1
        characterCreate1ViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

        val sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val backBtn = binding.backBtn
        val forwBtn = binding.forwBtn
        val characterView = binding.characterView
        val nameCharacter = binding.NameCharacter.editText
        val sexCharacter = binding.SexCharacter.editText
        val weightCharacter = binding.WeightCharacter.editText
        val heightCharacter = binding.HeightCharacter.editText



        backBtn.setOnClickListener {
            (requireActivity() as MainActivity).showConfirmationDialog(
                "Подтверждение возврата",
                "Данные не сохранены. Вы уверены, что хотите вернуться?",
                {
                    val activity = requireActivity() as MainActivity
                    val navView = activity.getNavView()
                    val menu = navView.menu
                    val menuItem1 = menu.findItem(R.id.navigation_info)
                    val menuItem2 = menu.findItem(R.id.navigation_character)
                    val menuItem3 = menu.findItem(R.id.navigation_lobby)
                    val menuItem4 = menu.findItem(R.id.navigation_account)
                    menuItem1.isVisible = true
                    menuItem2.isVisible = true
                    menuItem3.isVisible = true
                    menuItem4.isVisible = true
                    activity.navController.navigate(R.id.action_characterCreate1_to_navigation_character)
                },
                {  }
            )
        }


        forwBtn.setOnClickListener{
            GlobalScope.launch(Dispatchers.Main) {

                var token: String = ""
                sharedViewModel.token.observe(viewLifecycleOwner, Observer {
                    // updating data in displayMsg
                    token = it
                })


                //println(getCookieValue())

                val client = OkHttpClient.Builder().addInterceptor { chain ->
                    val original = chain.request()
                    val authorized = original.newBuilder()
                        .header("Authorization", token)
                        .build()
                    chain.proceed(authorized)
                }
                    .build()

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/class/getClasses")
                    .get()
                    .build()

                try {
                    val response = withContext(Dispatchers.IO) {
                        client.newCall(request).execute()
                    }

                    if (!response.isSuccessful) {
                        throw IOException("Запрос к серверу не был успешен:" +
                                " ${response.code} ${response.message}")
                    }

                    println("${response.code} ${response.message}")

                    val deserializedClassList = JsonHelper.deserializeListFromJsonCharacterClassList(response.body!!.string())
                    println("Десериализованный объект:\n$deserializedClassList")

                    sharedViewModel.characterClasses.value = deserializedClassList

                    val newCharacter = sharedViewModel.newCharacter.value
a@
                    newCharacter?.charName = nameCharacter?.text.toString()
                    newCharacter?.sex = sexCharacter?.text.toString().toBoolean()
                    newCharacter?.weight = weightCharacter?.text.toString().toInt()
                    newCharacter?.height = heightCharacter?.text.toString().toInt()

                    (activity as MainActivity).navController.navigate(R.id.action_characterCreate1_to_characterCreate2)


                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }
            }


        }
    }

}