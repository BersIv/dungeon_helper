package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyExpBinding
import com.example.dungeon_helper.shared.SharedViewModel
import okhttp3.WebSocket

class MasterLobbyExp : Fragment() {

    companion object {
        fun newInstance() = MasterLobbyExp()
    }

    private lateinit var viewModel: MasterLobbyExpViewModel
    private var _binding: FragmentMasterLobbyExpBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
       val masterLobbyExpViewModel = ViewModelProvider(this)[MasterLobbyExpViewModel::class.java]
        _binding = FragmentMasterLobbyExpBinding.inflate(inflater, container, false)
        val root: View = binding.root
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

        val sharedViewModel = ViewModelProvider(this)[SharedViewModel::class.java]

        var socket: WebSocket? = sharedViewModel.socket.value
        sharedViewModel.socket.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            socket = it
        })

        var id: String? = sharedViewModel.idToChange.value
        sharedViewModel.idToChange.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            id = it
        })

        val exp = binding.hp.editText

        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyExp_to_masterLobbyActChoose)
        }

        val acceptBtn = binding.acceptBtn
        acceptBtn.setOnClickListener {
            val expString = exp?.text

            socket?.send("change_exp $id $expString")

            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyExp_to_masterLobby)
        }
    }


}